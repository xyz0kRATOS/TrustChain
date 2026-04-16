package models

import (
	"time"
)

// CampaignStatus mirrors the DB status values for campaigns.
type CampaignStatus string

const (
	CampaignStatusPending   CampaignStatus = "PENDING"
	CampaignStatusActive    CampaignStatus = "ACTIVE"
	CampaignStatusCompleted CampaignStatus = "COMPLETED"
	CampaignStatusFrozen    CampaignStatus = "FROZEN"
	CampaignStatusRejected  CampaignStatus = "REJECTED"
)

// MilestoneStatus mirrors the DB status values for milestones.
type MilestoneStatus string

const (
	MilestoneStatusPending     MilestoneStatus = "PENDING"
	MilestoneStatusSubmitted   MilestoneStatus = "SUBMITTED"
	MilestoneStatusUnderReview MilestoneStatus = "UNDER_REVIEW"
	MilestoneStatusApproved    MilestoneStatus = "APPROVED"
	MilestoneStatusReleasing   MilestoneStatus = "RELEASING"
	MilestoneStatusCompleted   MilestoneStatus = "COMPLETED"
	MilestoneStatusOverdue     MilestoneStatus = "OVERDUE"
	MilestoneStatusDisputed    MilestoneStatus = "DISPUTED"
)

// Campaign is the DB model for the campaigns table.
type Campaign struct {
	ID                string         `db:"id"                  json:"id"`
	ContractAddress   *string        `db:"contract_address"    json:"contractAddress,omitempty"`
	CreatorWallet     string         `db:"creator_wallet"      json:"creatorWallet"`
	Category          *string        `db:"category"            json:"category,omitempty"`
	Name              string         `db:"name"                json:"name"`
	Description       string         `db:"description"         json:"description"`
	GoalAmountUSD     string         `db:"goal_amount_usd"     json:"goalAmountUsd"`
	GoalAmountWei     string         `db:"goal_amount_wei"     json:"goalAmountWei"` // NUMERIC(78) → string
	Status            CampaignStatus `db:"status"              json:"status"`
	DocumentHash      *string        `db:"document_hash"       json:"documentHash,omitempty"`
	ImageURL          *string        `db:"image_url"           json:"imageUrl,omitempty"`
	CreatorName       *string        `db:"creator_name"        json:"creatorName,omitempty"`
	CreatorEmail      *string        `db:"creator_email"       json:"creatorEmail,omitempty"`
	CreatorOrg        *string        `db:"creator_org"         json:"creatorOrg,omitempty"`
	CreatorCountry    *string        `db:"creator_country"     json:"creatorCountry,omitempty"`
	CreatorBio        *string        `db:"creator_bio"         json:"creatorBio,omitempty"`
	DocumentFileNames []string       `db:"document_file_names" json:"documentFileNames,omitempty"`
	IPFSEvidenceHash  *string        `db:"ipfs_evidence_hash"  json:"ipfsEvidenceHash,omitempty"`
	CreatedAt         time.Time      `db:"created_at"          json:"createdAt"`
	UpdatedAt         time.Time      `db:"updated_at"          json:"updatedAt"`
}

// Milestone is the DB model for the milestones table.
type Milestone struct {
	ID               string          `db:"id"                json:"id"`
	CampaignID       string          `db:"campaign_id"       json:"campaignId"`
	SequenceIndex    int             `db:"sequence_index"    json:"sequenceIndex"`
	Name             *string         `db:"name"              json:"name,omitempty"`
	Description      string          `db:"description"       json:"description"`
	AmountUSD        string          `db:"amount_usd"        json:"amountUsd"`
	AmountWei        string          `db:"amount_wei"        json:"amountWei"`
	Deadline         time.Time       `db:"deadline"          json:"deadline"`
	RequiredEvidence string          `db:"required_evidence" json:"requiredEvidence"`
	Status           MilestoneStatus `db:"status"            json:"status"`
	SubmittedAt      *time.Time      `db:"submitted_at"      json:"submittedAt,omitempty"`
	ApprovedAt       *time.Time      `db:"approved_at"       json:"approvedAt,omitempty"`
	ReleasedAt       *time.Time      `db:"released_at"       json:"releasedAt,omitempty"`
}

// Donation is the DB model for the donations table.
type Donation struct {
	ID              string    `db:"id"                json:"id"`
	CampaignID      string    `db:"campaign_id"       json:"campaignId"`
	DonorWallet     string    `db:"donor_wallet"      json:"donorWallet"`
	AmountWei       string    `db:"amount_wei"        json:"amountWei"`
	TxHash          string    `db:"tx_hash"           json:"txHash"`
	NFTTokenID      *int      `db:"nft_token_id"      json:"nftTokenId,omitempty"`
	IPFSMetadataCID *string   `db:"ipfs_metadata_cid" json:"ipfsMetadataCid,omitempty"`
	BlockNumber     int64     `db:"block_number"      json:"blockNumber"`
	DonatedAt       time.Time `db:"donated_at"        json:"donatedAt"`
}

// AccessGrant is the DB model for the access_grants table.
type AccessGrant struct {
	CampaignID  string    `db:"campaign_id"  json:"campaignId"`
	DonorWallet string    `db:"donor_wallet" json:"donorWallet"`
	GrantedAt   time.Time `db:"granted_at"   json:"grantedAt"`
}

// CampaignUpdate is the DB model for the campaign_updates table.
type CampaignUpdate struct {
	ID            string    `db:"id"             json:"id"`
	CampaignID    string    `db:"campaign_id"    json:"campaignId"`
	CreatorWallet string    `db:"creator_wallet" json:"creatorWallet"`
	Title         string    `db:"title"          json:"title"`
	Body          string    `db:"body"           json:"body"`
	UpdateType    string    `db:"update_type"    json:"updateType"`
	IPFSHashes    []string  `db:"ipfs_hashes"    json:"ipfsHashes"`
	UpdateHash    string    `db:"update_hash"    json:"updateHash"`
	OnChainTx     *string   `db:"on_chain_tx"    json:"onChainTx,omitempty"`
	MilestoneID   *string   `db:"milestone_id"   json:"milestoneId,omitempty"`
	IsFlagged     bool      `db:"is_flagged"     json:"isFlagged"`
	CreatedAt     time.Time `db:"created_at"     json:"createdAt"`
}

// AdminAction is the DB model for the admin_actions table.
type AdminAction struct {
	ID          string    `db:"id"           json:"id"`
	ActionType  string    `db:"action_type"  json:"actionType"`
	CampaignID  *string   `db:"campaign_id"  json:"campaignId,omitempty"`
	AdminWallet string    `db:"admin_wallet" json:"adminWallet"`
	Notes       *string   `db:"notes"        json:"notes,omitempty"`
	SafeTxHash  *string   `db:"safe_tx_hash" json:"safeTxHash,omitempty"`
	PerformedAt time.Time `db:"performed_at" json:"performedAt"`
}

// EmailQueue is the DB model for the email_queue table.
type EmailQueue struct {
	ID             string     `db:"id"              json:"id"`
	RecipientEmail string     `db:"recipient_email" json:"recipientEmail"`
	Subject        string     `db:"subject"         json:"subject"`
	TemplateName   string     `db:"template_name"   json:"templateName"`
	TemplateData   []byte     `db:"template_data"   json:"templateData"` // JSONB
	Status         string     `db:"status"          json:"status"`
	Attempts       int        `db:"attempts"        json:"attempts"`
	CreatedAt      time.Time  `db:"created_at"      json:"createdAt"`
	SentAt         *time.Time `db:"sent_at"         json:"sentAt,omitempty"`
}

// APIResponse is the standard envelope for all backend responses.
type APIResponse[T any] struct {
	Data  T       `json:"data"`
	Error *string `json:"error"`
	Meta  *Meta   `json:"meta,omitempty"`
}

// CampaignResponse is the API shape used by public and admin campaign endpoints.
type CampaignResponse struct {
	ID                string              `json:"id"`
	ContractAddress   *string             `json:"contractAddress"`
	CreatorWallet     string              `json:"creatorWallet"`
	Category          *string             `json:"category,omitempty"`
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	GoalAmountUSD     string              `json:"goalAmountUsd"`
	GoalAmountWei     *string             `json:"goalAmountWei"`
	Status            string              `json:"status"`
	DocumentHash      *string             `json:"documentHash"`
	ImageURL          *string             `json:"imageUrl"`
	DonorCount        int                 `json:"donorCount"`
	TotalRaisedWei    string              `json:"totalRaisedWei"`
	Milestones        []MilestoneResponse `json:"milestones"`
	CreatedAt         string              `json:"createdAt"`
	CreatorName       *string             `json:"creatorName,omitempty"`
	CreatorEmail      *string             `json:"creatorEmail,omitempty"`
	CreatorOrg        *string             `json:"creatorOrg,omitempty"`
	CreatorCountry    *string             `json:"creatorCountry,omitempty"`
	CreatorBio        *string             `json:"creatorBio,omitempty"`
	DocumentFileNames []string            `json:"documentFileNames,omitempty"`
}

type MilestoneResponse struct {
	ID               string  `json:"id"`
	SequenceIndex    int     `json:"sequenceIndex"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	AmountUSD        string  `json:"amountUsd"`
	AmountWei        *string `json:"amountWei"`
	Deadline         string  `json:"deadline"`
	RequiredEvidence string  `json:"requiredEvidence"`
	Status           string  `json:"status"`
}

type CampaignApplyRequest struct {
	Name              string                   `json:"name"`
	Category          string                   `json:"category"`
	Description       string                   `json:"description"`
	GoalAmountUSD     float64                  `json:"goalAmountUsd"`
	ImageURL          *string                  `json:"imageUrl"`
	CreatorWallet     string                   `json:"creatorWallet"`
	CreatorName       string                   `json:"creatorName"`
	CreatorEmail      string                   `json:"creatorEmail"`
	CreatorOrg        *string                  `json:"creatorOrg"`
	CreatorCountry    string                   `json:"creatorCountry"`
	CreatorBio        string                   `json:"creatorBio"`
	Milestones        []CampaignApplyMilestone `json:"milestones"`
	DocumentFileNames []string                 `json:"documentFileNames"`
}

type CampaignApplyMilestone struct {
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	AmountUSD        float64 `json:"amountUsd"`
	Deadline         string  `json:"deadline"`
	RequiredEvidence string  `json:"requiredEvidence"`
}

type CampaignApplyResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type ActivityEvent struct {
	ID           string  `json:"id"`
	Type         string  `json:"type"`
	CampaignName string  `json:"campaignName"`
	CampaignID   string  `json:"campaignId"`
	Amount       *string `json:"amount"`
	Wallet       *string `json:"wallet"`
	Timestamp    string  `json:"timestamp"`
	TxHash       *string `json:"txHash"`
}

// Meta carries pagination info.
type Meta struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// APIError returns a response with a non-nil error string and nil data.
func NewErrorResponse(msg string) APIResponse[any] {
	return APIResponse[any]{Error: &msg}
}
