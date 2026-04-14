// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/governance/TimelockController.sol";

// Interfaces — Campaign talks to these contracts but does not inherit them
interface IReputationRegistry {
    function updateScore(address creatorWallet, int256 delta, string calldata reason) external;
}

interface IDonationNFT {
    function mint(address donor, uint256 amount, string calldata ipfsMetadataHash) external returns (uint256);
}

/**
 * @title Campaign
 * @notice One instance of this contract exists per TrustChain campaign.
 *         Holds donated ETH and enforces milestone-gated, Timelock-delayed withdrawals.
 *
 * Lifecycle:
 *   ACTIVE     → Accepting donations, waiting for milestone submissions
 *   COMPLETED  → All milestones released, campaign closed successfully
 *   FROZEN     → Admin froze the campaign; all remaining funds proportionally refunded
 *
 * Milestone release flow:
 *   1. Creator submits evidence off-chain (backend + IPFS)
 *   2. Admin multi-sig calls approveMilestone() → schedules release in TimelockController
 *   3. 48 hours pass
 *   4. Anyone calls executeMilestoneRelease() → ETH sent to creator, reputation updated
 *
 * Security:
 *   - ReentrancyGuard on all ETH-moving functions
 *   - Checks-Effects-Interactions pattern throughout
 *   - Milestones locked at construction — no setters exist
 *   - Only admin multi-sig can approve milestones or freeze
 *   - Only TimelockController can execute fund releases
 */
contract Campaign is ReentrancyGuard {

    // ── Enums ─────────────────────────────────────────────────────────────────

    enum CampaignStatus { ACTIVE, COMPLETED, FROZEN }

    enum MilestoneStatus {
        PENDING,
        SUBMITTED,
        UNDER_REVIEW,
        APPROVED,
        RELEASING,
        COMPLETED,
        OVERDUE,
        DISPUTED
    }

    // ── Structs ───────────────────────────────────────────────────────────────

    struct Milestone {
        string  description;
        uint256 amount;           // ETH in wei to release on completion
        uint256 deadline;         // Unix timestamp
        string  requiredEvidence; // Short string describing what proof is needed
        MilestoneStatus status;
        bool    fundsReleased;
    }

    // ── Immutable state (set at construction, never changed) ──────────────────

    address public immutable creator;
    address public immutable adminMultiSig;
    address public immutable timelockController;
    address public immutable reputationRegistry;
    address public immutable donationNFT;

    uint256 public immutable goalAmount;      // Total fundraising goal in wei
    uint256 public immutable campaignId;      // Unique ID from CampaignFactory
    bytes32 public immutable documentHash;    // SHA-256 hash of verified documents

    // ── Mutable state ─────────────────────────────────────────────────────────

    CampaignStatus public status;
    uint256 public totalRaised;
    uint256 public totalReleased;

    Milestone[] public milestones;

    /// @dev donor address → total amount donated in wei
    mapping(address => uint256) public donations;

    /// @dev Ordered list of donors (for proportional refund iteration on freeze)
    address[] public donorList;

    /// @dev Tracks whether a donor has been added to donorList already
    mapping(address => bool) private _isDonor;

    // ── Events ────────────────────────────────────────────────────────────────

    event DonationReceived(
        address indexed donor,
        uint256 amount,
        uint256 indexed campaignId,
        uint256 timestamp
    );

    event MilestoneApproved(
        uint256 indexed milestoneIndex,
        address indexed approver,
        uint256 timestamp
    );

    event FundsReleased(
        uint256 indexed milestoneIndex,
        uint256 amount,
        address indexed recipient
    );

    event CampaignFrozen(address indexed frozenBy, uint256 timestamp);

    event DonorRefunded(address indexed donor, uint256 amount);

    event CampaignCompleted(uint256 timestamp);

    event CampaignUpdated(
        uint256 indexed campaignId,
        bytes32 updateHash,
        string  updateType,
        uint256 timestamp
    );

    event MilestoneAmended(
        uint256 indexed milestoneIndex,
        bytes32 amendmentHash,
        uint256 timestamp
    );

    // ── Modifiers ─────────────────────────────────────────────────────────────

    modifier onlyAdmin() {
        require(msg.sender == adminMultiSig, "Campaign: caller is not admin");
        _;
    }

    modifier onlyTimelock() {
        require(msg.sender == timelockController, "Campaign: caller is not timelock");
        _;
    }

    modifier onlyActive() {
        require(status == CampaignStatus.ACTIVE, "Campaign: not active");
        _;
    }

    modifier validMilestoneIndex(uint256 index) {
        require(index < milestones.length, "Campaign: invalid milestone index");
        _;
    }

    // ── Constructor ───────────────────────────────────────────────────────────

    /**
     * @param _creator              Wallet address of the campaign creator.
     * @param _adminMultiSig        Safe multi-sig address with approval authority.
     * @param _timelockController   OpenZeppelin TimelockController address (48h delay).
     * @param _reputationRegistry   ReputationRegistry contract address.
     * @param _donationNFT          DonationNFT contract address.
     * @param _goalAmount           Total fundraising goal in wei.
     * @param _campaignId           Unique campaign ID assigned by CampaignFactory.
     * @param _documentHash         SHA-256 hash of verified campaign documents.
     * @param _milestoneDescriptions  Array of milestone description strings.
     * @param _milestoneAmounts       Array of ETH amounts (wei) per milestone.
     * @param _milestoneDeadlines     Array of Unix timestamps (deadlines).
     * @param _milestoneEvidence      Array of required evidence type strings.
     */
    constructor(
        address _creator,
        address _adminMultiSig,
        address _timelockController,
        address _reputationRegistry,
        address _donationNFT,
        uint256 _goalAmount,
        uint256 _campaignId,
        bytes32 _documentHash,
        string[]  memory _milestoneDescriptions,
        uint256[] memory _milestoneAmounts,
        uint256[] memory _milestoneDeadlines,
        string[]  memory _milestoneEvidence
    ) {
        require(_creator       != address(0), "Campaign: zero creator");
        require(_adminMultiSig != address(0), "Campaign: zero admin");
        require(
            _milestoneDescriptions.length == _milestoneAmounts.length &&
            _milestoneAmounts.length      == _milestoneDeadlines.length &&
            _milestoneDeadlines.length    == _milestoneEvidence.length,
            "Campaign: milestone array length mismatch"
        );
        require(_milestoneDescriptions.length > 0, "Campaign: no milestones");

        // Validate milestone amounts sum to goal
        uint256 sum;
        for (uint256 i = 0; i < _milestoneAmounts.length; i++) {
            sum += _milestoneAmounts[i];
        }
        require(sum == _goalAmount, "Campaign: milestone amounts must sum to goal");

        creator             = _creator;
        adminMultiSig       = _adminMultiSig;
        timelockController  = _timelockController;
        reputationRegistry  = _reputationRegistry;
        donationNFT         = _donationNFT;
        goalAmount          = _goalAmount;
        campaignId          = _campaignId;
        documentHash        = _documentHash;
        status              = CampaignStatus.ACTIVE;

        // Lock milestones permanently at construction
        for (uint256 i = 0; i < _milestoneDescriptions.length; i++) {
            milestones.push(Milestone({
                description:      _milestoneDescriptions[i],
                amount:           _milestoneAmounts[i],
                deadline:         _milestoneDeadlines[i],
                requiredEvidence: _milestoneEvidence[i],
                status:           MilestoneStatus.PENDING,
                fundsReleased:    false
            }));
        }
    }

    // ── Donation ──────────────────────────────────────────────────────────────

    /**
     * @notice Accept an ETH donation.
     *         Automatically mints a soulbound NFT receipt to the donor.
     *         Minimum donation: enforced off-chain by the frontend (suggested $5 equiv).
     * @param ipfsMetadataHash  IPFS CID for this donation's NFT metadata — passed by frontend.
     */
    function donate(string calldata ipfsMetadataHash) external payable onlyActive nonReentrant {
        require(msg.value > 0, "Campaign: donation must be > 0");

        // Effects — update state before external calls
        totalRaised += msg.value;
        donations[msg.sender] += msg.value;

        if (!_isDonor[msg.sender]) {
            _isDonor[msg.sender] = true;
            donorList.push(msg.sender);
        }

        emit DonationReceived(msg.sender, msg.value, campaignId, block.timestamp);

        // Interactions — external call after state is updated
        IDonationNFT(donationNFT).mint(msg.sender, msg.value, ipfsMetadataHash);
    }

    // ── Milestone approval (admin only) ───────────────────────────────────────

    /**
     * @notice Admin multi-sig approves a milestone and schedules its fund release
     *         through the TimelockController (48-hour delay).
     *         Enforces sequential milestone completion — index N requires N-1 to be COMPLETED.
     * @param milestoneIndex  Zero-based index of the milestone to approve.
     */
    function approveMilestone(uint256 milestoneIndex)
        external
        onlyAdmin
        onlyActive
        validMilestoneIndex(milestoneIndex)
    {
        // Enforce sequential unlock — milestone 0 must complete before milestone 1, etc.
        if (milestoneIndex > 0) {
            require(
                milestones[milestoneIndex - 1].status == MilestoneStatus.COMPLETED,
                "Campaign: previous milestone not completed"
            );
        }

        Milestone storage m = milestones[milestoneIndex];
        require(
            m.status == MilestoneStatus.UNDER_REVIEW || m.status == MilestoneStatus.SUBMITTED,
            "Campaign: milestone not ready for approval"
        );

        // Effects
        m.status = MilestoneStatus.APPROVED;
        emit MilestoneApproved(milestoneIndex, msg.sender, block.timestamp);

        // Interactions — schedule release through Timelock
        // The Timelock will call executeMilestoneRelease() after its delay
        bytes memory callData = abi.encodeWithSelector(
            this.executeMilestoneRelease.selector,
            milestoneIndex
        );

        TimelockController(payable(timelockController)).schedule(
            address(this),   // target
            0,               // value
            callData,        // data
            bytes32(0),      // predecessor (none)
            bytes32(milestoneIndex), // salt (unique per milestone)
            48 hours         // delay
        );
    }

    /**
     * @notice Execute a milestone's fund release after the Timelock delay has passed.
     *         Can be called by anyone once the Timelock has matured — no signature needed.
     *         In practice, your backend or keeper triggers this automatically.
     * @param milestoneIndex  Zero-based index of the milestone being released.
     */
    function executeMilestoneRelease(uint256 milestoneIndex)
        external
        onlyTimelock
        nonReentrant
        validMilestoneIndex(milestoneIndex)
    {
        Milestone storage m = milestones[milestoneIndex];
        require(m.status == MilestoneStatus.APPROVED, "Campaign: milestone not approved");
        require(!m.fundsReleased, "Campaign: funds already released");
        require(address(this).balance >= m.amount, "Campaign: insufficient contract balance");

        // Checks-Effects-Interactions
        m.status       = MilestoneStatus.COMPLETED;
        m.fundsReleased = true;
        totalReleased  += m.amount;

        emit FundsReleased(milestoneIndex, m.amount, creator);

        // Update reputation for on-time milestone
        IReputationRegistry(reputationRegistry).updateScore(
            creator,
            10, // MILESTONE_COMPLETED constant from ReputationRegistry
            "Milestone completed"
        );

        // Check if all milestones are now complete
        bool allComplete = true;
        for (uint256 i = 0; i < milestones.length; i++) {
            if (milestones[i].status != MilestoneStatus.COMPLETED) {
                allComplete = false;
                break;
            }
        }

        if (allComplete) {
            status = CampaignStatus.COMPLETED;
            emit CampaignCompleted(block.timestamp);
            IReputationRegistry(reputationRegistry).updateScore(
                creator,
                25, // CAMPAIGN_COMPLETED constant
                "Campaign fully completed"
            );
        }

        // Transfer ETH to creator — after all state changes
        (bool sent, ) = creator.call{value: m.amount}("");
        require(sent, "Campaign: ETH transfer to creator failed");
    }

    // ── Milestone status updates (admin only, off-chain sync) ─────────────────

    /**
     * @notice Admin updates a milestone's status to reflect off-chain state changes.
     *         Used to sync SUBMITTED / UNDER_REVIEW / OVERDUE / DISPUTED states.
     * @param milestoneIndex  Zero-based index of the milestone.
     * @param newStatus       New status enum value.
     */
    function setMilestoneStatus(uint256 milestoneIndex, MilestoneStatus newStatus)
        external
        onlyAdmin
        validMilestoneIndex(milestoneIndex)
    {
        milestones[milestoneIndex].status = newStatus;
    }

    // ── Freeze + proportional refund (admin only) ─────────────────────────────

    /**
     * @notice Freeze the campaign and proportionally refund all donors
     *         from the remaining (unreleased) contract balance.
     *         Triggered by: creator disappears, fraud detected, deadline passed, creator requests cancel.
     * @dev    Uses the checks-effects-interactions pattern.
     *         Iterates donorList — safe for v1 (max 10 campaigns, limited donors).
     *         For v2 with large donor lists, consider a pull-based refund pattern instead.
     */
    function freeze() external onlyAdmin onlyActive nonReentrant {
        // Effects first
        status = CampaignStatus.FROZEN;
        emit CampaignFrozen(msg.sender, block.timestamp);

        IReputationRegistry(reputationRegistry).updateScore(
            creator,
            -20, // CAMPAIGN_FROZEN constant
            "Campaign frozen"
        );

        uint256 remaining = address(this).balance;
        if (remaining == 0 || totalRaised == 0) return;

        // Calculate and send each donor's proportional share
        // Using totalRaised as the denominator ensures correct proportions
        // even if some funds were already released to the creator
        for (uint256 i = 0; i < donorList.length; i++) {
            address donor = donorList[i];
            uint256 donated = donations[donor];
            if (donated == 0) continue;

            // Proportional share of remaining balance
            uint256 refundAmount = (donated * remaining) / totalRaised;

            if (refundAmount > 0) {
                // Effects — zero out before sending
                donations[donor] = 0;

                emit DonorRefunded(donor, refundAmount);

                // Interactions
                (bool sent, ) = donor.call{value: refundAmount}("");
                // If a donor's wallet rejects ETH (e.g. a contract), we do not revert —
                // we skip them. They can pursue recovery off-chain.
                // This prevents one bad actor from blocking all other refunds.
                if (!sent) {
                    donations[donor] = refundAmount; // Restore for manual recovery
                }
            }
        }
    }

    // ── On-chain update anchoring ─────────────────────────────────────────────

    /**
     * @notice Anchor a campaign update hash on-chain.
     *         The full update content lives in PostgreSQL + IPFS.
     *         This call just records the tamper-evident SHA-256 hash.
     * @param updateHash  SHA-256 hash of the full update record from the backend.
     * @param updateType  One of: progress_update, milestone_evidence, setback_report,
     *                    amendment_request, general_announcement.
     */
    function recordUpdate(bytes32 updateHash, string calldata updateType)
        external
        onlyAdmin
    {
        emit CampaignUpdated(campaignId, updateHash, updateType, block.timestamp);
    }

    /**
     * @notice Anchor a milestone amendment hash on-chain.
     *         The amendment content and original milestone remain visible.
     *         This creates an immutable addendum record — does not modify any milestone struct.
     * @param milestoneIndex  Which milestone was amended.
     * @param amendmentHash   SHA-256 hash of (original milestone + amendment details + timestamp).
     */
    function recordAmendment(uint256 milestoneIndex, bytes32 amendmentHash)
        external
        onlyAdmin
        validMilestoneIndex(milestoneIndex)
    {
        emit MilestoneAmended(milestoneIndex, amendmentHash, block.timestamp);
    }

    // ── View helpers ──────────────────────────────────────────────────────────

    function getMilestoneCount() external view returns (uint256) {
        return milestones.length;
    }

    function getMilestone(uint256 index) external view returns (Milestone memory) {
        require(index < milestones.length, "Campaign: invalid index");
        return milestones[index];
    }

    function getDonorCount() external view returns (uint256) {
        return donorList.length;
    }

    function getContractBalance() external view returns (uint256) {
        return address(this).balance;
    }

    // ── Fallback ──────────────────────────────────────────────────────────────

    /// @dev Reject plain ETH transfers — donors must use donate() to get their NFT
    receive() external payable {
        revert("Campaign: use donate() to contribute");
    }
}
