// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title DonationNFT
 * @notice Soulbound (non-transferable) ERC-721 NFT receipt minted to a donor's
 *         wallet every time they donate to any TrustChain Campaign.
 *
 * Soulbound enforcement:
 *   All transfer functions (_update override) revert if the sender is not
 *   the zero address — meaning only minting (from == zero) is allowed.
 *   Approvals are also blocked for completeness.
 *
 * Metadata:
 *   Token URI points to IPFS. Each token's metadata JSON contains:
 *     - campaign name, campaign ID, donor address, amount donated, timestamp
 *   The IPFS hash is passed in at mint time by the Campaign contract.
 *
 * Access control:
 *   Only authorised Campaign contracts (registered by CampaignFactory/owner)
 *   can call mint(). Public cannot mint directly.
 */
contract DonationNFT is ERC721, Ownable {

    // ── Storage ──────────────────────────────────────────────────────────────

    uint256 private _nextTokenId;

    /// @dev Maps token ID → IPFS metadata URI
    mapping(uint256 => string) private _tokenURIs;

    /// @dev Tracks which Campaign contracts are authorised to call mint()
    mapping(address => bool) public authorisedCampaigns;

    // ── Events ────────────────────────────────────────────────────────────────

    event DonationNFTMinted(
        uint256 indexed tokenId,
        address indexed donor,
        address indexed campaignContract,
        uint256 amount,
        string ipfsMetadataHash
    );

    event CampaignAuthorised(address indexed campaignAddress);

    // ── Modifiers ─────────────────────────────────────────────────────────────

    modifier onlyAuthorisedCampaign() {
        require(
            authorisedCampaigns[msg.sender],
            "DonationNFT: caller is not an authorised Campaign"
        );
        _;
    }

    // ── Constructor ───────────────────────────────────────────────────────────

    constructor() ERC721("TrustChain Donation Receipt", "TCDR") Ownable(msg.sender) {}

    // ── Admin functions ───────────────────────────────────────────────────────

    /**
     * @notice Register a Campaign contract as authorised to mint NFTs.
     *         Called by CampaignFactory immediately after deploying each Campaign.
     * @param campaignAddress The address of the newly deployed Campaign contract.
     */
    function authoriseCampaign(address campaignAddress) external onlyOwner {
        require(campaignAddress != address(0), "DonationNFT: zero address");
        authorisedCampaigns[campaignAddress] = true;
        emit CampaignAuthorised(campaignAddress);
    }

    // ── Mint (Campaign contracts only) ────────────────────────────────────────

    /**
     * @notice Mint a soulbound donation receipt NFT to a donor.
     *         Called automatically by Campaign.donate() on every successful donation.
     * @param donor             Wallet address receiving the NFT.
     * @param amount            ETH amount donated (in wei) — stored in metadata.
     * @param ipfsMetadataHash  IPFS CID for this donation's metadata JSON.
     */
    function mint(
        address donor,
        uint256 amount,
        string calldata ipfsMetadataHash
    ) external onlyAuthorisedCampaign returns (uint256 tokenId) {
        require(donor != address(0), "DonationNFT: zero address donor");

        tokenId = _nextTokenId;
        _nextTokenId++;

        _safeMint(donor, tokenId);
        _tokenURIs[tokenId] = string(abi.encodePacked("ipfs://", ipfsMetadataHash));

        emit DonationNFTMinted(tokenId, donor, msg.sender, amount, ipfsMetadataHash);
    }

    // ── Soulbound enforcement ─────────────────────────────────────────────────

    /**
     * @dev Override ERC721 _update to block all transfers.
     *      Minting (from == address(0)) is permitted.
     *      Any transfer or burn is blocked.
     */
    function _update(
        address to,
        uint256 tokenId,
        address auth
    ) internal override returns (address) {
        address from = _ownerOf(tokenId);
        require(from == address(0), "DonationNFT: token is soulbound and cannot be transferred");
        return super._update(to, tokenId, auth);
    }

    /**
     * @dev Block all approvals — soulbound tokens cannot be approved for transfer.
     */
    function approve(address, uint256) public pure override {
        revert("DonationNFT: approvals disabled for soulbound tokens");
    }

    function setApprovalForAll(address, bool) public pure override {
        revert("DonationNFT: approvals disabled for soulbound tokens");
    }

    // ── Metadata ──────────────────────────────────────────────────────────────

    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        require(_ownerOf(tokenId) != address(0), "DonationNFT: token does not exist");
        return _tokenURIs[tokenId];
    }

    function totalSupply() external view returns (uint256) {
        return _nextTokenId;
    }
}
