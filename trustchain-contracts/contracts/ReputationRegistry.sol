// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title ReputationRegistry
 * @notice Permanent public scoreboard for campaign creator wallets.
 *         Only registered Campaign contracts can write scores.
 *         Anyone can read any wallet's score.
 *
 * Score rules (enforced by Campaign.sol on lifecycle events):
 *   +10  Milestone completed on time
 *   +25  Campaign fully completed
 *    -5  Milestone missed deadline
 *   -20  Campaign frozen or disputed
 *
 * Security:
 *   - Only the CampaignFactory (owner) can register authorised Campaign addresses
 *   - Only registered Campaign addresses can call updateScore()
 *   - Score floor is 0 — cannot go negative
 */
contract ReputationRegistry is Ownable {

    // ── Storage ──────────────────────────────────────────────────────────────

    /// @dev Maps a creator wallet address to their cumulative reputation score
    mapping(address => int256) private _scores;

    /// @dev Tracks which contract addresses are authorised Campaign contracts
    mapping(address => bool) public authorisedCampaigns;

    // ── Events ────────────────────────────────────────────────────────────────

    event ReputationUpdated(
        address indexed wallet,
        int256 newScore,
        string reason
    );

    event CampaignAuthorised(address indexed campaignAddress);

    // ── Score change constants ────────────────────────────────────────────────

    int256 public constant MILESTONE_COMPLETED   =  10;
    int256 public constant CAMPAIGN_COMPLETED    =  25;
    int256 public constant MILESTONE_MISSED      =  -5;
    int256 public constant CAMPAIGN_FROZEN       = -20;

    // ── Modifiers ─────────────────────────────────────────────────────────────

    modifier onlyAuthorisedCampaign() {
        require(
            authorisedCampaigns[msg.sender],
            "ReputationRegistry: caller is not an authorised Campaign"
        );
        _;
    }

    // ── Constructor ───────────────────────────────────────────────────────────

    constructor() Ownable(msg.sender) {}

    // ── Admin functions (CampaignFactory calls these) ─────────────────────────

    /**
     * @notice Register a newly deployed Campaign contract as authorised to write scores.
     *         Called by CampaignFactory immediately after deploying each Campaign.
     * @param campaignAddress The address of the newly deployed Campaign contract.
     */
    function authoriseCampaign(address campaignAddress) external onlyOwner {
        require(campaignAddress != address(0), "ReputationRegistry: zero address");
        authorisedCampaigns[campaignAddress] = true;
        emit CampaignAuthorised(campaignAddress);
    }

    // ── Write functions (only Campaign contracts) ─────────────────────────────

    /**
     * @notice Update a creator's reputation score.
     *         Only callable by an authorised Campaign contract.
     * @param creatorWallet  The wallet address of the campaign creator.
     * @param delta          Positive or negative score change (use the constants above).
     * @param reason         Short human-readable reason string — indexed by The Graph.
     */
    function updateScore(
        address creatorWallet,
        int256 delta,
        string calldata reason
    ) external onlyAuthorisedCampaign {
        require(creatorWallet != address(0), "ReputationRegistry: zero address");

        int256 current = _scores[creatorWallet];
        int256 updated = current + delta;

        // Floor at zero — reputation cannot go negative
        if (updated < 0) {
            updated = 0;
        }

        _scores[creatorWallet] = updated;

        emit ReputationUpdated(creatorWallet, updated, reason);
    }

    // ── Read functions (public) ───────────────────────────────────────────────

    /**
     * @notice Returns the current reputation score for any wallet.
     *         Fully public — no access restriction.
     * @param wallet  The wallet address to query.
     */
    function getScore(address wallet) external view returns (int256) {
        return _scores[wallet];
    }
}
