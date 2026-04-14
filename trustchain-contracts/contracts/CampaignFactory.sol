// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./Campaign.sol";
import "./ReputationRegistry.sol";
import "./DonationNFT.sol";
import "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * @title CampaignFactory
 * @notice The single entry point for creating TrustChain campaigns.
 *         Only the admin multi-sig can create campaigns — not the public.
 *         Deploys a new Campaign contract per campaign, registers it with
 *         ReputationRegistry and DonationNFT, and maintains a full registry.
 *
 * This contract is owned by the admin multi-sig (Safe wallet).
 * Ownership can be transferred if the multi-sig address changes.
 *
 * Deployment order:
 *   1. Deploy ReputationRegistry
 *   2. Deploy DonationNFT
 *   3. Deploy CampaignFactory (passing above addresses + TimelockController address)
 *   4. Transfer ownership of ReputationRegistry to CampaignFactory
 *   5. Transfer ownership of DonationNFT to CampaignFactory
 *
 * After step 4-5, only CampaignFactory can authorise new campaigns in the
 * registry contracts — keeping the authorisation flow centralised here.
 */
contract CampaignFactory is Ownable {

    // ── Storage ──────────────────────────────────────────────────────────────

    /// @dev Ordered list of all deployed Campaign contract addresses
    address[] public allCampaigns;

    /// @dev Maps campaignId → Campaign contract address
    mapping(uint256 => address) public campaignById;

    /// @dev Internal counter for unique campaign IDs
    uint256 private _nextCampaignId;

    // ── Immutable references to sibling contracts ─────────────────────────────

    address public immutable reputationRegistry;
    address public immutable donationNFT;
    address public immutable timelockController;

    // ── Events ────────────────────────────────────────────────────────────────

    event CampaignCreated(
        address indexed campaignAddress,
        address indexed creatorWallet,
        uint256 indexed campaignId,
        uint256 goalAmount,
        uint256 milestoneCount,
        bytes32 documentHash
    );

    // ── Constructor ───────────────────────────────────────────────────────────

    /**
     * @param _adminMultiSig       Safe multi-sig that will own this factory and all campaigns.
     * @param _reputationRegistry  Deployed ReputationRegistry address.
     * @param _donationNFT         Deployed DonationNFT address.
     * @param _timelockController  Deployed TimelockController address (48h delay).
     */
    constructor(
        address _adminMultiSig,
        address _reputationRegistry,
        address _donationNFT,
        address _timelockController
    ) Ownable(_adminMultiSig) {
        require(_reputationRegistry  != address(0), "Factory: zero registry");
        require(_donationNFT         != address(0), "Factory: zero nft");
        require(_timelockController  != address(0), "Factory: zero timelock");

        reputationRegistry = _reputationRegistry;
        donationNFT        = _donationNFT;
        timelockController = _timelockController;
    }

    // ── Campaign creation (admin multi-sig only) ──────────────────────────────

    /**
     * @notice Deploy a new Campaign contract for an approved campaign.
     *         Called by the admin multi-sig only — not by campaign creators directly.
     *
     * @param creatorWallet           Wallet address of the verified campaign creator.
     * @param goalAmount              Total fundraising goal in wei.
     * @param documentHash            SHA-256 hash of all verified campaign documents.
     * @param milestoneDescriptions   Ordered array of milestone description strings.
     * @param milestoneAmounts        Ordered array of ETH amounts (wei) per milestone.
     *                                Must sum exactly to goalAmount.
     * @param milestoneDeadlines      Ordered array of Unix timestamps for each deadline.
     * @param milestoneEvidence       Ordered array of required evidence type strings.
     *
     * @return campaignAddress  The address of the newly deployed Campaign contract.
     */
    function createCampaign(
        address creatorWallet,
        uint256 goalAmount,
        bytes32 documentHash,
        string[]  calldata milestoneDescriptions,
        uint256[] calldata milestoneAmounts,
        uint256[] calldata milestoneDeadlines,
        string[]  calldata milestoneEvidence
    ) external onlyOwner returns (address campaignAddress) {
        require(creatorWallet != address(0), "Factory: zero creator");
        require(goalAmount > 0, "Factory: goal must be > 0");
        require(milestoneDescriptions.length > 0, "Factory: no milestones");

        uint256 newCampaignId = _nextCampaignId;
        _nextCampaignId++;

        // Deploy Campaign — passing all immutable parameters at construction
        Campaign campaign = new Campaign(
            creatorWallet,
            owner(),             // adminMultiSig = this factory's owner
            timelockController,
            reputationRegistry,
            donationNFT,
            goalAmount,
            newCampaignId,
            documentHash,
            milestoneDescriptions,
            milestoneAmounts,
            milestoneDeadlines,
            milestoneEvidence
        );

        campaignAddress = address(campaign);

        // Register in factory's own registry
        allCampaigns.push(campaignAddress);
        campaignById[newCampaignId] = campaignAddress;

        // Authorise this campaign in ReputationRegistry and DonationNFT
        // (Factory must own both contracts for these calls to succeed)
        ReputationRegistry(reputationRegistry).authoriseCampaign(campaignAddress);
        DonationNFT(donationNFT).authoriseCampaign(campaignAddress);
        
        // Grant the Campaign the PROPOSER_ROLE so it can schedule milestone releases
        // (Factory must be the admin of the TimelockController for this to succeed)
        TimelockController payableTimelock = TimelockController(payable(timelockController));
        payableTimelock.grantRole(payableTimelock.PROPOSER_ROLE(), campaignAddress);

        emit CampaignCreated(
            campaignAddress,
            creatorWallet,
            newCampaignId,
            goalAmount,
            milestoneDescriptions.length,
            documentHash
        );
    }

    // ── Registry reads (public) ───────────────────────────────────────────────

    /**
     * @notice Returns the total number of campaigns ever created.
     */
    function getCampaignCount() external view returns (uint256) {
        return allCampaigns.length;
    }

    /**
     * @notice Returns a paginated slice of all campaign addresses.
     *         Use this to avoid fetching the entire array in one call.
     * @param offset  Start index (0-based).
     * @param limit   Number of addresses to return.
     */
    function getCampaigns(uint256 offset, uint256 limit)
        external
        view
        returns (address[] memory result)
    {
        uint256 total = allCampaigns.length;
        if (offset >= total) return new address[](0);

        uint256 end = offset + limit;
        if (end > total) end = total;

        result = new address[](end - offset);
        for (uint256 i = offset; i < end; i++) {
            result[i - offset] = allCampaigns[i];
        }
    }
}
