import { Address, BigInt, Bytes } from "@graphprotocol/graph-ts"

import {
  CampaignCreated
} from "../generated/CampaignFactory/CampaignFactory"

import {
  DonationNFTMinted
} from "../generated/DonationNFT/DonationNFT"

import {
  ReputationUpdated
} from "../generated/ReputationRegistry/ReputationRegistry"

import {
  CampaignCompleted,
  CampaignFrozen,
  CampaignUpdated,
  DonationReceived,
  DonorRefunded,
  FundsReleased,
  MilestoneAmended,
  MilestoneApproved
} from "../generated/templates/CampaignTemplate/Campaign"

import { CampaignTemplate } from "../generated/templates"

import {
  Campaign,
  Donation,
  DonorRefund,
  FundRelease,
  FundraiserProfile,
  Milestone
} from "../generated/schema"

function zero(): BigInt {
  return BigInt.zero()
}

function getOrCreateProfile(wallet: Address): FundraiserProfile {
  let id = wallet.toHexString()
  let profile = FundraiserProfile.load(id)
  if (profile == null) {
    profile = new FundraiserProfile(id)
    profile.wallet = wallet
    profile.score = zero()
    profile.save()
  }
  return profile
}

function getOrCreateMilestone(campaignId: string, index: i32): Milestone {
  let id = campaignId + "-" + index.toString()
  let milestone = Milestone.load(id)
  if (milestone == null) {
    milestone = new Milestone(id)
    milestone.campaign = campaignId
    milestone.index = index
    milestone.status = "PENDING"
    milestone.amount = zero()
    milestone.deadline = zero()
    milestone.fundsReleased = false
    milestone.save()
  }
  return milestone as Milestone
}

function donationId(txHash: Bytes, donor: Address, campaign: Address): string {
  return txHash.toHexString() + "-" + donor.toHexString() + "-" + campaign.toHexString()
}

export function handleCampaignCreated(event: CampaignCreated): void {
  let campaignAddress = event.params.campaignAddress
  let id = campaignAddress.toHexString()

  let campaign = Campaign.load(id)
  if (campaign != null) {
    CampaignTemplate.create(campaignAddress)
    return
  }

  let creator = getOrCreateProfile(event.params.creatorWallet)

  campaign = new Campaign(id)
  campaign.creator = creator.id
  campaign.goalAmount = event.params.goalAmount
  campaign.status = "ACTIVE"
  campaign.totalRaised = zero()
  campaign.createdAt = event.block.timestamp
  campaign.save()

  let count = event.params.milestoneCount.toI32()
  for (let i = 0; i < count; i++) {
    let milestone = new Milestone(id + "-" + i.toString())
    milestone.campaign = id
    milestone.index = i
    milestone.status = "PENDING"
    milestone.amount = zero()
    milestone.deadline = zero()
    milestone.fundsReleased = false
    milestone.save()
  }

  CampaignTemplate.create(campaignAddress)
}

export function handleDonationReceived(event: DonationReceived): void {
  let campaignId = event.address.toHexString()
  let campaign = Campaign.load(campaignId)
  if (campaign == null) {
    return
  }

  let id = donationId(event.transaction.hash, event.params.donor, event.address)
  let donation = Donation.load(id)
  if (donation == null) {
    donation = new Donation(id)
  }

  donation.campaign = campaignId
  donation.donor = event.params.donor
  donation.amount = event.params.amount
  donation.timestamp = event.params.timestamp
  donation.save()

  campaign.totalRaised = campaign.totalRaised.plus(event.params.amount)
  campaign.save()
}

export function handleMilestoneApproved(event: MilestoneApproved): void {
  let campaignId = event.address.toHexString()
  let campaign = Campaign.load(campaignId)
  if (campaign == null) {
    return
  }

  let milestone = getOrCreateMilestone(campaignId, event.params.milestoneIndex.toI32())
  milestone.status = "APPROVED"
  milestone.save()
}

export function handleFundsReleased(event: FundsReleased): void {
  let campaignId = event.address.toHexString()
  let campaign = Campaign.load(campaignId)
  if (campaign == null) {
    return
  }

  let milestoneIndex = event.params.milestoneIndex.toI32()
  let milestone = getOrCreateMilestone(campaignId, milestoneIndex)
  milestone.status = "COMPLETED"
  milestone.fundsReleased = true
  if (milestone.amount.equals(zero())) {
    milestone.amount = event.params.amount
  }
  milestone.save()

  let release = new FundRelease(event.transaction.hash.toHexString() + "-" + event.logIndex.toString())
  release.campaign = campaignId
  release.milestone = milestone.id
  release.amount = event.params.amount
  release.timestamp = event.block.timestamp
  release.save()
}

export function handleCampaignFrozen(event: CampaignFrozen): void {
  let campaign = Campaign.load(event.address.toHexString())
  if (campaign == null) {
    return
  }

  campaign.status = "FROZEN"
  campaign.save()
}

export function handleDonorRefunded(event: DonorRefunded): void {
  let campaignId = event.address.toHexString()
  let campaign = Campaign.load(campaignId)
  if (campaign == null) {
    return
  }

  let refund = new DonorRefund(event.transaction.hash.toHexString() + "-" + event.logIndex.toString())
  refund.campaign = campaignId
  refund.donor = event.params.donor
  refund.amount = event.params.amount
  refund.timestamp = event.block.timestamp
  refund.save()
}

export function handleCampaignCompleted(event: CampaignCompleted): void {
  let campaign = Campaign.load(event.address.toHexString())
  if (campaign == null) {
    return
  }

  campaign.status = "COMPLETED"
  campaign.save()
}

export function handleReputationUpdated(event: ReputationUpdated): void {
  let profile = getOrCreateProfile(event.params.wallet)
  profile.score = event.params.newScore
  profile.save()
}

export function handleDonationNFTMinted(event: DonationNFTMinted): void {
  let id = donationId(
    event.transaction.hash,
    event.params.donor,
    event.params.campaignContract
  )

  let donation = Donation.load(id)
  if (donation == null) {
    return
  }

  donation.nftTokenId = event.params.tokenId
  donation.save()
}

export function handleCampaignUpdated(event: CampaignUpdated): void {
  let campaign = Campaign.load(event.address.toHexString())
  if (campaign == null) {
    return
  }

  if (campaign.status == "") {
    campaign.status = "ACTIVE"
    campaign.save()
  }
}

export function handleMilestoneAmended(event: MilestoneAmended): void {
  let campaignId = event.address.toHexString()
  let campaign = Campaign.load(campaignId)
  if (campaign == null) {
    return
  }

  let milestone = getOrCreateMilestone(campaignId, event.params.milestoneIndex.toI32())
  milestone.status = "AMENDED"
  milestone.save()
}
