import assert from "node:assert/strict";
import { describe, it } from "node:test";
import { network } from "hardhat";
import { parseEther, keccak256, stringToHex, getAddress } from "viem";

// Add BigInt serialization for assertions
BigInt.prototype["toJSON"] = function () {
  return this.toString();
};

describe("TrustChain Ecosystem", async function () {
  const { viem } = await network.connect();

  let publicClient: any;
  let adminMultiSig: any;
  let creator: any;
  let donor1: any;
  let donor2: any;
  
  let reputationRegistry: any;
  let donationNFT: any;
  let timelock: any;
  let factory: any;
  let campaign: any;

  it("Should setup environment and accounts", async function () {
    publicClient = await viem.getPublicClient();
    const wallets = await viem.getWalletClients();
    
    adminMultiSig = wallets[0];
    creator = wallets[1];
    donor1 = wallets[2];
    donor2 = wallets[3];
  });

  it("Should deploy core infrastructure", async function () {
    // 1. Deploy ReputationRegistry
    reputationRegistry = await viem.deployContract("ReputationRegistry");
    
    // 2. Deploy DonationNFT
    donationNFT = await viem.deployContract("DonationNFT");
    
    // 3. Deploy TestTimelockController (1 hour delay for tests)
    const minDelay = 3600n; // 1 hr
    timelock = await viem.deployContract("TestTimelockController", [
      minDelay,
      [adminMultiSig.account.address], // proposer
      [adminMultiSig.account.address], // executor
      adminMultiSig.account.address    // admin
    ]);
    
    // 4. Deploy CampaignFactory
    factory = await viem.deployContract("CampaignFactory", [
      adminMultiSig.account.address,
      reputationRegistry.address,
      donationNFT.address,
      timelock.address
    ]);

    // 5. Transfer ownership and grant roles
    await reputationRegistry.write.transferOwnership([factory.address]);
    await donationNFT.write.transferOwnership([factory.address]);
    
    const DEFAULT_ADMIN_ROLE = await timelock.read.DEFAULT_ADMIN_ROLE();
    const grantTx = await timelock.write.grantRole([DEFAULT_ADMIN_ROLE, factory.address]);
    await publicClient.waitForTransactionReceipt({ hash: grantTx });
    
    // verify the factory has the admin role
    const hasAdmin = await timelock.read.hasRole([DEFAULT_ADMIN_ROLE, factory.address]);
    console.log("Factory has DEFAULT_ADMIN_ROLE:", hasAdmin);
    
    // Assert owners are now factory
    assert.equal((await reputationRegistry.read.owner()).toLowerCase(), factory.address.toLowerCase());
    assert.equal((await donationNFT.read.owner()).toLowerCase(), factory.address.toLowerCase());
  });

  it("Should create a campaign via Factory", async function () {
    const goalAmount = parseEther("10"); // 10 ETH
    
    const milestoneDescriptions = ["Initial Setup", "Mid-point", "Final Delivery"];
    const milestoneAmounts = [
      parseEther("2"),
      parseEther("3"),
      parseEther("5")
    ];
    // Deadlines
    const now = Math.floor(Date.now() / 1000);
    const milestoneDeadlines = [
      BigInt(now + 86400 * 7),
      BigInt(now + 86400 * 14),
      BigInt(now + 86400 * 30)
    ];
    const milestoneEvidence = ["Receipts", "Images and logs", "Final report"];
    const documentHash = keccak256(stringToHex("test documents"));

    // Expected next ID is 0
    const nextId = await factory.read.getCampaignCount();

    const txHash = await factory.write.createCampaign([
      creator.account.address,
      goalAmount,
      documentHash,
      milestoneDescriptions,
      milestoneAmounts,
      milestoneDeadlines,
      milestoneEvidence
    ]);

    const receipt = await publicClient.waitForTransactionReceipt({ hash: txHash });

    // Verify it was deployed and registered
    const campaignCount = await factory.read.getCampaignCount();
    assert.equal(campaignCount, nextId + 1n);

    const campaignAddress = await factory.read.campaignById([nextId]);
    console.log("Campaign address deployed:", campaignAddress);
    
    // Instantiate Campaign
    campaign = await viem.getContractAt("Campaign", campaignAddress);
    
    const PROPOSER_ROLE = await timelock.read.PROPOSER_ROLE();
    const hasRole = await timelock.read.hasRole([PROPOSER_ROLE, campaignAddress]);
    console.log("Campaign has PROPOSER_ROLE:", hasRole);
    
    // Ensure creator is set properly
    assert.equal((await campaign.read.creator()).toLowerCase(), creator.account.address.toLowerCase());
  });

  it("Should allow a donor to donate and receive NFT", async function () {
    const donationAmount = parseEther("1");
    const ipfsMetadataHash = "QmTestHash1234";
    
    const tx = await campaign.write.donate([ipfsMetadataHash], {
      value: donationAmount,
      account: donor1.account,
    });
    
    await publicClient.waitForTransactionReceipt({ hash: tx });
    
    // Campaign balance should be 1 ETH
    const balance = await publicClient.getBalance({ address: campaign.address });
    assert.equal(balance, donationAmount);

    // NFT balances
    const nftBalance = await donationNFT.read.balanceOf([donor1.account.address]);
    assert.equal(nftBalance, 1n);
    
    // Token URI test
    const uri = await donationNFT.read.tokenURI([0n]);
    assert.equal(uri, "ipfs://QmTestHash1234");
  });

  it("Should approve milestone 0", async function () {
    // Current milestone 0 is pending
    // We update status to SUBMITTED first using setMilestoneStatus
    // setMilestoneStatus takes (uint256, MilestoneStatus enum values)
    // MilestoneStatus: PENDING(0), SUBMITTED(1), UNDER_REVIEW(2), APPROVED(3), RELEASING(4), COMPLETED(5), OVERDUE(6), DISPUTED(7)
    
    await campaign.write.setMilestoneStatus([0n, 1]); // SUBMITTED
    
    // Approve milestone
    const tx = await campaign.write.approveMilestone([0n]);
    // The TimelockController schedule is called internally.
    const receipt = await publicClient.waitForTransactionReceipt({ hash: tx });
    assert.equal(receipt.status, "success");
    
    const milestone = await campaign.read.milestones([0n]);
    // Status should be APPROVED = 3
    assert.equal(milestone[4], 3); // index of status is 4 in the struct
  });

});
