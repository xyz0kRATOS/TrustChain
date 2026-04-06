# TrustChain — AI Coding Assistant Project Rules
# Paste this as your project rules file in Cursor / Windsurf / any AI coding IDE.
# Every file you generate must follow these rules exactly.

## Project Identity
You are building TrustChain — a blockchain-based donations and fund tracking 
platform on Base (Ethereum L2). This is a production-grade application. 
Every decision prioritises security, auditability, and transparency.

---

## Tech Stack — Use These. Do Not Substitute.

### Blockchain
- Network: Base (L2). Testnet: Base Sepolia. Mainnet: Base mainnet.
- Contracts: Solidity 0.8.20, OpenZeppelin, Hardhat
- Web3 library: viem + wagmi v2
- Wallet connection: RainbowKit v2
- Contract interaction: generated TypeScript ABIs via wagmi CLI

### Frontend
- Framework: Next.js 14 (App Router)
- Language: TypeScript (strict mode, no any types allowed)
- Styling: Tailwind CSS
- Graph visualisation: Cytoscape.js (fund flow graph)
- Data fetching: TanStack Query (React Query)
- State management: Zustand for global UI state only
- Forms: React Hook Form + Zod validation

### Backend (Go) — Frontend calls these endpoints
- Base URL stored in: NEXT_PUBLIC_API_URL env variable
- Auth: wallet signature verification (SIWE — Sign In With Ethereum)
- All API responses follow: { data: T, error: string | null }

### Indexing
- The Graph Protocol — GraphQL endpoint in: NEXT_PUBLIC_GRAPH_URL env variable
- Use Apollo Client for Graph queries

### File Storage
- IPFS via Pinata — backend handles pinning, frontend receives CID back

---

## Environment Variables
Never hardcode addresses or keys. Always use env variables.
Required .env.local variables:

NEXT_PUBLIC_CHAIN_ID=84532                          # Base Sepolia testnet
NEXT_PUBLIC_FACTORY_ADDRESS=0x...                   # CampaignFactory contract
NEXT_PUBLIC_REGISTRY_ADDRESS=0x...                  # ReputationRegistry contract  
NEXT_PUBLIC_NFT_ADDRESS=0x...                       # DonationNFT contract
NEXT_PUBLIC_API_URL=http://localhost:8080            # Go backend
NEXT_PUBLIC_GRAPH_URL=https://api.thegraph.com/...  # The Graph subgraph endpoint
NEXT_PUBLIC_WALLET_CONNECT_ID=...                   # WalletConnect project ID

---

## Project Folder Structure
Maintain this structure exactly. Do not reorganise.

trustchain-web/
  app/                          # Next.js App Router pages
    (public)/                   # No auth required
      page.tsx                  # Homepage
      campaigns/
        page.tsx                # Campaign browser
        [id]/
          page.tsx              # Campaign detail
      fundraiser/[wallet]/
        page.tsx                # Public fundraiser profile
    (protected)/                # Requires wallet connection
      donor/
        page.tsx                # Donor profile + NFT receipts
      apply/
        page.tsx                # Campaign application form
    admin/                      # Admin only — multi-sig wallet gated
      page.tsx                  # Dashboard
      campaigns/[id]/
        page.tsx                # Campaign review + approval
  components/
    blockchain/                 # All Web3 components
      WalletButton.tsx
      DonateModal.tsx
      MilestoneApproval.tsx
      FreezeButton.tsx
    campaign/                   # Campaign UI components
      CampaignCard.tsx
      MilestoneTracker.tsx
      DonorList.tsx
      UpdateFeed.tsx
      FundFlowGraph.tsx         # Cytoscape.js visualisation
    layout/
      Navbar.tsx
      Footer.tsx
      LiveActivityFeed.tsx      # WebSocket-powered sidebar
    ui/                         # Base UI components (no blockchain logic)
      Button.tsx
      Badge.tsx
      StatusPill.tsx
      ProgressBar.tsx
  hooks/
    useWallet.ts                # Wallet connection state
    useCampaign.ts              # Fetch campaign data
    useDonate.ts                # Donation transaction flow
    useAdmin.ts                 # Admin-gated actions
    useWebSocket.ts             # Live activity feed
  lib/
    wagmi.ts                    # wagmi config + chain setup
    apollo.ts                   # Apollo Client for The Graph
    contracts.ts                # Contract addresses + ABIs
    graph-queries.ts            # All GraphQL query definitions
    api.ts                      # Go backend API client
    siwe.ts                     # Sign In With Ethereum helpers
  types/
    campaign.ts                 # Campaign, Milestone, Donor types
    graph.ts                    # GraphQL response types
    api.ts                      # Backend API response types

---

## Smart Contract Addresses and ABIs

Store ALL contract interactions in lib/contracts.ts.
Never import ABIs directly in components.

The four contracts and their key functions:

### CampaignFactory
Functions frontend calls:
- getCampaignCount() → uint256
- getCampaigns(offset, limit) → address[]
- campaignById(id) → address
Events to listen for:
- CampaignCreated(campaignAddress, creatorWallet, campaignId, goalAmount, milestoneCount, documentHash)

### Campaign (one address per campaign)
Functions frontend calls:
- donate(ipfsMetadataHash: string) → payable
- getMilestone(index) → Milestone struct
- getMilestoneCount() → uint256
- getDonorCount() → uint256
- getContractBalance() → uint256
- status() → 0=ACTIVE, 1=COMPLETED, 2=FROZEN
- totalRaised() → uint256
- creator() → address
- goalAmount() → uint256
- campaignId() → uint256
Events:
- DonationReceived(donor, amount, campaignId, timestamp)
- MilestoneApproved(milestoneIndex, approver, timestamp)
- FundsReleased(milestoneIndex, amount, recipient)
- CampaignFrozen(frozenBy, timestamp)
- DonorRefunded(donor, amount)

### ReputationRegistry
Functions frontend calls:
- getScore(wallet: address) → int256

### DonationNFT
Functions frontend calls:
- tokenURI(tokenId) → string
- totalSupply() → uint256

---

## Wallet Connection Rules

Use RainbowKit + wagmi v2. Configure in lib/wagmi.ts.
- Support: MetaMask, Coinbase Wallet, WalletConnect
- Restrict to Base chain only — show error if user is on wrong network
- Auto-prompt network switch if user connects on wrong chain
- Never ask for wallet signature unless the action requires it
- Admin actions (approveMilestone, freeze) require multi-sig — 
  frontend shows "This action requires Safe multi-sig" and links 
  to the pending Safe transaction

---

## The Donation Flow (Most Critical Path)

This is the highest-risk user journey. Build it with zero ambiguity.

Step 1: User clicks Donate on campaign page
Step 2: DonateModal opens — amount input + confirmation
Step 3: Frontend calls backend POST /api/campaigns/:id/prepare-donation
         → Backend pins NFT metadata to IPFS via Pinata
         → Backend returns: { ipfsCid: string, estimatedGas: number }
Step 4: Frontend calls donate(ipfsCid) on Campaign contract via wagmi
         → Pass msg.value = donation amount in wei
Step 5: Show transaction pending state (spinner + tx hash link to Basescan)
Step 6: Wait for receipt (1 confirmation on Base is sufficient)
Step 7: On success: show NFT receipt preview, update campaign total
Step 8: On failure: show specific error (insufficient funds, wrong network, etc.)

NEVER:
- Submit a donation without the IPFS CID — NFT won't have metadata
- Show success before transaction is confirmed on-chain
- Let a user donate 0 ETH
- Skip the confirmation step

---

## The Fund Flow Graph (Cytoscape.js)

Build in components/campaign/FundFlowGraph.tsx

Data source: Go backend GET /api/campaigns/:id/graph
Returns: { nodes: Node[], edges: Edge[] }

Node types:
- donor: blue circle, size proportional to total donated
- contract: dark circle (the Campaign contract itself)  
- milestone: green circle (completed), amber (pending), red (frozen)
- creator: teal circle

Edge rules:
- Donation: donor → contract (label: amount + date)
- Release: contract → milestone → creator (label: amount released)
- Refund: contract → donor (label: refund amount, red colour)

Public view: donor addresses truncated (0x4a2f...8c3d)
Verified donor view: full addresses visible

Re-render trigger: new DonationReceived or FundsReleased events via WebSocket

---

## Access Control (4 Tiers)

Implement in a useAccessLevel() hook.

Tier 0 — Public (no wallet):
  Can see: campaign name, total raised, milestone names, fund flow graph 
  (truncated addresses), fundraiser reputation score

Tier 1 — Verified Donor (wallet connected + donated to this campaign):
  Check: backend GET /api/campaigns/:id/access?wallet=0x...
  Can see: everything above + full donor list with addresses and amounts

Tier 2 — Campaign Creator (wallet matches campaign.creator):
  Can see: all donor data + milestone submission tools + evidence upload

Tier 3 — Admin (wallet matches adminMultiSig):
  Can see: everything + admin controls (approve, freeze, review queue)
  Note: approval actions go through Safe multi-sig, not direct tx

Gating pattern:
  Use <AccessGate requiredTier={1} campaignId={id}> wrapper component.
  Show upgrade prompt if tier is insufficient (e.g. "Donate to see full donor list")

---

## The Graph — GraphQL Queries

All queries live in lib/graph-queries.ts
Use Apollo Client. Never fetch blockchain data directly from components.

Key entities in the subgraph:
- Campaign: id, creatorWallet, status, goalAmount, totalRaised, milestones[]
- Donation: id, donor, amount, campaignId, timestamp, nftTokenId
- Milestone: id, campaignId, index, status, amount, deadline, fundsReleased
- FundRelease: id, milestoneIndex, amount, recipient, timestamp
- FundraiserProfile: wallet, reputationScore, totalCampaigns, completedCampaigns
- DonationReceipt: tokenId, donor, campaignId, amount, ipfsMetadataHash

---

## Live Activity Feed (WebSocket)

Build in components/layout/LiveActivityFeed.tsx
Hook: hooks/useWebSocket.ts

Connect to: NEXT_PUBLIC_API_URL/ws
The Go backend pushes JSON events when it detects on-chain activity.

Event shape:
{
  type: "donation" | "milestone_approved" | "funds_released" | "campaign_frozen" | "campaign_created",
  campaignId: string,
  campaignName: string,
  amount?: string,
  wallet?: string,
  milestoneIndex?: number,
  timestamp: string
}

Display format:
- donation: "0x4a2f...8c3d donated 0.05 ETH to Clean Water Nairobi — 5s ago"
- milestone_approved: "Milestone 2 approved for School Rebuild — releasing in 48h"
- funds_released: "0.8 ETH released to School Rebuild Kenya"

Show last 20 events. New events animate in from the top.
Auto-reconnect if WebSocket drops.

---

## TypeScript Types (types/campaign.ts)

export type CampaignStatus = "ACTIVE" | "COMPLETED" | "FROZEN"

export type MilestoneStatus = 
  "PENDING" | "SUBMITTED" | "UNDER_REVIEW" | "APPROVED" | 
  "RELEASING" | "COMPLETED" | "OVERDUE" | "DISPUTED"

export interface Milestone {
  index: number
  description: string
  amount: bigint          // wei
  deadline: number        // unix timestamp
  requiredEvidence: string
  status: MilestoneStatus
  fundsReleased: boolean
}

export interface Campaign {
  id: string
  contractAddress: string
  creator: string
  status: CampaignStatus
  goalAmount: bigint
  totalRaised: bigint
  totalReleased: bigint
  milestones: Milestone[]
  donorCount: number
  documentHash: string
  name: string            // from backend PostgreSQL, not on-chain
  description: string     // from backend PostgreSQL
  imageUrl?: string       // from backend PostgreSQL
}

export interface Donor {
  wallet: string
  amount: bigint
  timestamp: number
  nftTokenId: number
}

---

## Error Handling Rules

Always handle these specific blockchain errors with human-readable messages:

- User rejected transaction → "You cancelled the transaction."
- Insufficient funds → "Your wallet doesn't have enough ETH. You need X ETH plus gas."
- Wrong network → "Please switch to Base network in your wallet."
- Contract revert: "Campaign not active" → "This campaign is no longer accepting donations."
- Contract revert: "use donate()" → "Please use the Donate button — do not send ETH directly."
- Network timeout → "Transaction is taking longer than expected. Check Basescan for status: [link]"
- Never show raw error objects or hex revert data to users.

---

## Milestone Status Display

Map contract enum values to UI display:

PENDING        → grey pill  "Awaiting submission"
SUBMITTED      → blue pill  "Evidence submitted"
UNDER_REVIEW   → amber pill "Under review"
APPROVED       → teal pill  "Approved — releasing in Xh Xm"  (countdown timer)
RELEASING      → teal pill  "Releasing..."
COMPLETED      → green pill "Completed" + green checkmark
OVERDUE        → red pill   "Overdue"
DISPUTED       → red pill   "Disputed"

---

## Code Quality Rules

- TypeScript strict mode. No `any`. No `as unknown as X`.
- Every component has a clearly typed Props interface.
- No inline styles — Tailwind classes only.
- No useEffect for data fetching — use TanStack Query.
- Every blockchain-reading hook must handle: loading, error, and empty states.
- Every transaction must show: pending → confirmed → success/failure states.
- Accessibility: all interactive elements have aria labels.
- No console.log in production code — use a logger utility.
- All monetary amounts stored and computed as bigint (wei). 
  Format for display only using viem's formatEther().

---

## Basescan Links

Always link to the right network.
Testnet: https://sepolia.basescan.org/tx/{hash}
Mainnet: https://basescan.org/tx/{hash}

Build a helper: lib/basescan.ts
  getTransactionUrl(hash: string): string
  getAddressUrl(address: string): string
  getTokenUrl(contractAddress: string, tokenId: number): string

Use NEXT_PUBLIC_CHAIN_ID to determine which URL to use.

---

## What to Build First (Follow This Order)

1. lib/wagmi.ts — wallet config
2. lib/contracts.ts — addresses + ABIs
3. components/blockchain/WalletButton.tsx — connect button
4. app/(public)/campaigns/[id]/page.tsx — campaign detail page (read only first)
5. components/blockchain/DonateModal.tsx — the donation flow
6. hooks/useDonate.ts — transaction logic
7. components/campaign/MilestoneTracker.tsx — milestone progress display
8. components/campaign/FundFlowGraph.tsx — Cytoscape.js graph
9. components/layout/LiveActivityFeed.tsx — WebSocket feed
10. app/(public)/campaigns/page.tsx — campaign browser
11. app/(public)/page.tsx — homepage
12. app/(protected)/donor/page.tsx — donor profile + NFT receipts
13. app/admin/ — admin dashboard (build last)

Build each item completely (with loading, error, empty states) 
before moving to the next. Do not scaffold everything at once.