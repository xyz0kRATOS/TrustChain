# TrustChain E2E Test Guide (Base Sepolia)

All testing in this guide uses Base Sepolia (`chainId=84532`) with faucet ETH only.

## Step 0 - Setup (One Time)

- Get test ETH from https://faucet.quicknode.com/base/sepolia for:
  - Deployer wallet (0.1 ETH)
  - Admin wallet (0.05 ETH)
  - Donor wallet 1 (0.05 ETH)
  - Donor wallet 2 (0.05 ETH)
- Deploy contracts:
  - `cd trustchain-contracts`
  - `npx hardhat run scripts/deploy.ts --network baseSepolia`
- Copy `trustchain-contracts/deployed-addresses.json` values into:
  - `trustchain-web/.env.local`
  - `trustchain-backend/.env`
- Seed local DB:
  - `cd trustchain-backend`
  - `go run scripts/seed.go --reset`
- Start backend:
  - `go run cmd/server/main.go`
- Start frontend:
  - `cd ../trustchain-web`
  - `npm run dev`
- Open:
  - `http://localhost:3000`

## Step 1 - Creator Applies

Action:
- Connect a non-admin test wallet
- Open `/apply`
- Submit:
  - Name: `E2E Test Campaign`
  - Goal: `$100`
  - Milestone 1: `First milestone`, `$60`, 30 days from now, `Invoice and delivery photo`
  - Milestone 2: `Second milestone`, `$40`, 60 days from now, `Final report`

Expected:
- Success state with real application ID

Verify:
- `psql $DATABASE_URL -c "SELECT id, name, status FROM campaigns ORDER BY created_at DESC LIMIT 1"`
- Most recent row has `status='pending'`

If it fails:
- Confirm backend is running on `:8080`
- Confirm `NEXT_PUBLIC_API_URL` points to backend
- Confirm wallet is connected before submit

## Step 2 - Admin Reviews and Approves

Action:
- Switch wallet to admin wallet
- Open `/admin`
- Approve pending application

Expected:
- Card leaves pending queue
- Success toast appears

Verify:
- `psql $DATABASE_URL -c "SELECT id, name, status FROM campaigns WHERE name='E2E Test Campaign'"`
- Row has `status='live'`

If it fails:
- Confirm `NEXT_PUBLIC_ADMIN_WALLET` in frontend env
- Confirm `ADMIN_WALLET` in backend env
- Confirm request includes `X-Admin-Wallet` header

## Step 3 - Campaign Appears Publicly

Action:
- Open `/campaigns`
- Click `E2E Test Campaign`

Expected:
- Detail page loads with no 404
- Shows campaign name, description, milestones
- Shows contract as `Pending deployment`

If it fails:
- Re-run seed + apply flow
- Confirm `/api/campaigns` and `/api/campaigns/:id` return data

## Step 4 - Deploy Approved Campaign (Phase 2)

Action:
- Use `/api/admin/campaigns/:id/deploy`

Expected:
- Returns `txHash` and deployed `contractAddress`
- Campaign row updates with contract address

Note:
- Current codebase includes endpoint scaffolding; full on-chain deployment logic is still marked as a stub.

## Step 5 - Donor Makes Donation (Phase 2)

Action:
- Connect donor wallet 1
- Donate `0.01 ETH`

Expected:
- Wallet confirmation
- Tx confirmation
- NFT receipt minted message

Verify:
- Basescan tx exists
- `donations` table has latest row

## Step 6 - Second Donor (Phase 2)

Action:
- Connect donor wallet 2
- Donate `0.005 ETH`

Expected:
- Campaign raised amount updates

## Step 7 - Approve Milestone (Phase 2)

Action:
- Admin calls `/api/admin/campaigns/:id/milestones/:idx/approve`

Expected:
- Milestone status becomes approved with timelock countdown

## Step 8 - Execute Release After Timelock (Phase 2)

Action:
- Wait timelock delay
- Admin calls `/api/admin/campaigns/:id/milestones/:idx/execute-release`

Expected:
- Milestone status completed
- Creator receives ETH
- Reputation score updates

## Step 9 - Verify Fund Flow Graph (Phase 2)

Action:
- Open campaign detail graph

Expected:
- Donor -> campaign -> milestone -> creator flow

## Step 10 - Verify Live Activity Feed

Expected feed order (newest first):
- Funds released
- Milestone approved
- Donations
- Campaign approved/live
- Campaign applied
