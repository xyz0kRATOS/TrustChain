package models

// BlockchainEvent represents a generic on-chain event pushed via WebSocket.
type BlockchainEvent struct {
	Type           string  `json:"type"`
	CampaignID     string  `json:"campaignId"`
	CampaignName   string  `json:"campaignName"`
	Amount         *string `json:"amount,omitempty"`
	Wallet         *string `json:"wallet,omitempty"`
	MilestoneIndex *int    `json:"milestoneIndex,omitempty"`
	Timestamp      string  `json:"timestamp"`
}

// DonationReceivedEvent maps to the Campaign contract DonationReceived event.
type DonationReceivedEvent struct {
	Donor      string `json:"donor"`
	Amount     string `json:"amount"`     // wei as string
	CampaignID string `json:"campaignId"`
	TxHash     string `json:"txHash"`
	Timestamp  int64  `json:"timestamp"`
}

// MilestoneApprovedEvent maps to the Campaign contract MilestoneApproved event.
type MilestoneApprovedEvent struct {
	MilestoneIndex int    `json:"milestoneIndex"`
	Approver       string `json:"approver"`
	CampaignID     string `json:"campaignId"`
	TxHash         string `json:"txHash"`
	Timestamp      int64  `json:"timestamp"`
}

// FundsReleasedEvent maps to the Campaign contract FundsReleased event.
type FundsReleasedEvent struct {
	MilestoneIndex int    `json:"milestoneIndex"`
	Amount         string `json:"amount"` // wei as string
	Recipient      string `json:"recipient"`
	CampaignID     string `json:"campaignId"`
	TxHash         string `json:"txHash"`
	Timestamp      int64  `json:"timestamp"`
}

// CampaignFrozenEvent maps to the Campaign contract CampaignFrozen event.
type CampaignFrozenEvent struct {
	FrozenBy   string `json:"frozenBy"`
	CampaignID string `json:"campaignId"`
	TxHash     string `json:"txHash"`
	Timestamp  int64  `json:"timestamp"`
}

// CampaignCreatedEvent maps to the CampaignFactory CampaignCreated event.
type CampaignCreatedEvent struct {
	CampaignAddress string `json:"campaignAddress"`
	CreatorWallet   string `json:"creatorWallet"`
	CampaignID      string `json:"campaignId"`
	GoalAmount      string `json:"goalAmount"` // wei as string
	MilestoneCount  int    `json:"milestoneCount"`
	DocumentHash    string `json:"documentHash"`
	TxHash          string `json:"txHash"`
}
