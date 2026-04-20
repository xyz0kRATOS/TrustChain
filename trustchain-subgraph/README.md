# TrustChain Subgraph

This subgraph indexes TrustChain contracts on Base Sepolia:
- CampaignFactory
- Campaign
- DonationNFT
- ReputationRegistry

All contract addresses in subgraph.yaml are placeholders and must be replaced after deployment.

## Prerequisites

- Node.js 18+
- npm
- Graph CLI (installed through package.json scripts)
- Deployed contract addresses on Base Sepolia

## Setup

1. Install dependencies:
   npm install

2. Update placeholder addresses in subgraph.yaml:
   - CampaignFactory source.address
   - Campaign source.address (optional static source)
   - DonationNFT source.address
   - ReputationRegistry source.address

3. Generate types:
   npm run codegen

4. Build the subgraph:
   npm run build

## Deploy to The Graph Studio

1. Authenticate:
   graph auth --studio <DEPLOY_KEY>

2. Deploy:
   graph deploy --studio trustchain-subgraph

## Deploy to Local Graph Node

1. Start local graph-node + IPFS + PostgreSQL stack.

2. Create and deploy:
   npm run create-local
   npm run deploy-local

## Notes

- CampaignFactory handler automatically creates dynamic CampaignTemplate data sources for new campaign contracts.
- Event coverage includes:
  - CampaignCreated
  - DonationReceived
  - MilestoneApproved
  - FundsReleased
  - CampaignFrozen
  - DonorRefunded
  - CampaignCompleted
  - ReputationUpdated
  - DonationNFTMinted
  - CampaignUpdated
  - MilestoneAmended
