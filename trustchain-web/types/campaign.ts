export interface MilestoneResponse {
  id: string;
  sequenceIndex: number;
  name: string;
  description: string;
  amountUsd: string;
  amountWei: string | null;
  deadline: string;
  requiredEvidence: string;
  status: string;
}

export interface CampaignResponse {
  id: string;
  contractAddress: string | null;
  creatorWallet: string;
  category?: string | null;
  name: string;
  description: string;
  goalAmountUsd: string;
  goalAmountWei: string | null;
  status: string;
  documentHash: string | null;
  imageUrl: string | null;
  donorCount: number;
  totalRaisedWei: string;
  milestones: MilestoneResponse[];
  createdAt: string;
  creatorName?: string | null;
  creatorEmail?: string | null;
  creatorOrg?: string | null;
  creatorCountry?: string | null;
  creatorBio?: string | null;
  documentFileNames?: string[];
}

export interface ActivityResponse {
  id: string;
  type:
    | 'campaign_applied'
    | 'campaign_approved'
    | 'campaign_live'
    | 'donation'
    | 'milestone_approved'
    | 'funds_released';
  campaignName: string;
  campaignId: string;
  amount: string | null;
  wallet: string | null;
  timestamp: string;
  txHash: string | null;
}

export interface ApiEnvelope<T> {
  data: T;
  error: string | null;
}
