# TrustChain — Go Backend Rules Prompt
# Paste this as your project rules file in Cursor / Windsurf for the backend.

## Identity
You are building the Go backend for TrustChain — a blockchain donations platform.
The backend handles everything that cannot live on-chain: identity verification,
email notifications, file storage, access control, and serving the frontend.

---

## Tech Stack
Language:        Go 1.22+
Framework:       Gin (HTTP router)
Database:        PostgreSQL 15 via pgx/v5 driver
Migrations:      golang-migrate
Blockchain:      go-ethereum (geth) for contract interactions + event watching
IPFS:            Pinata REST API via standard net/http
Email:           Resend Go SDK
Auth:            SIWE (Sign In With Ethereum) + JWT
WebSocket:       gorilla/websocket
Graph queries:   standard net/http POST to The Graph GraphQL endpoint
Config:          viper + godotenv
Logging:         zerolog (structured JSON logs)
Testing:         standard testing package + testify

---

## Project Folder Structure
Maintain this exactly:

trustchain-backend/
  cmd/
    server/
      main.go                # Entry point only — wire everything up here

  internal/
    api/
      router.go              # Gin router + middleware registration

      middleware/
        auth.go              # JWT + SIWE verification middleware
        cors.go              # CORS for frontend origin
        ratelimit.go         # Basic rate limiting

      handlers/
        campaigns.go         # Campaign CRUD handlers
        donations.go         # Donation prepare + metadata handlers
        milestones.go        # Milestone submission + review handlers
        admin.go             # Admin dashboard handlers
        auth.go              # SIWE login/logout handlers
        updates.go           # Campaign update handlers
        graph.go             # Graph proxy + enrichment handlers

    blockchain/
      client.go              # go-ethereum client setup
      contracts.go           # Contract ABI bindings + call helpers
      watcher.go             # Event listener — watches all 4 contracts
      safe.go                # Safe multi-sig transaction construction

    db/
      postgres.go            # Database connection pool

      migrations/
        001_create_campaigns.sql
        002_create_milestones.sql
        003_create_donations.sql
        004_create_updates.sql
        005_create_access_grants.sql
        006_create_admin_actions.sql
        007_create_email_queue.sql

      queries/
        campaigns.go         # All campaign DB queries
        milestones.go
        donations.go
        updates.go
        access.go

    ipfs/
      pinata.go              # Pinata client — pin JSON and files

    email/
      resend.go              # Resend client + email templates
      templates/             # Email HTML templates

    graph/
      client.go              # The Graph GraphQL client
      queries.go             # All GraphQL query strings

    websocket/
      hub.go                 # WebSocket connection hub
      broadcaster.go         # Event → WebSocket message

    models/
      campaign.go            # Go structs for all DB entities
      blockchain.go          # Structs matching contract event types

    config/
      config.go              # Viper config loader

  .env                       # Never commit
  .gitignore
  go.mod
  go.sum

---

## Database Schema

### campaigns table
id                 UUID PRIMARY KEY
contract_address   VARCHAR(42) NULLABLE
creator_wallet     VARCHAR(42) NOT NULL
name               VARCHAR(255) NOT NULL
description        TEXT NOT NULL
goal_amount_wei    NUMERIC(78) NOT NULL
status             VARCHAR(20) NOT NULL
document_hash      VARCHAR(66) NULLABLE
image_url          TEXT NULLABLE
ipfs_evidence_hash TEXT NULLABLE
created_at         TIMESTAMP DEFAULT NOW()
updated_at         TIMESTAMP DEFAULT NOW()

---

### milestones table
id                 UUID PRIMARY KEY
campaign_id        UUID REFERENCES campaigns(id)
sequence_index     INT NOT NULL
description        TEXT NOT NULL
amount_wei         NUMERIC(78) NOT NULL
deadline           TIMESTAMP NOT NULL
required_evidence  TEXT NOT NULL
status             VARCHAR(20) NOT NULL
submitted_at       TIMESTAMP NULLABLE
approved_at        TIMESTAMP NULLABLE
released_at        TIMESTAMP NULLABLE

---

### donations table
id               UUID PRIMARY KEY
campaign_id      UUID REFERENCES campaigns(id)
donor_wallet     VARCHAR(42) NOT NULL
amount_wei       NUMERIC(78) NOT NULL
tx_hash          VARCHAR(66) NOT NULL
nft_token_id     INT NULLABLE
ipfs_metadata_cid TEXT NULLABLE
block_number     BIGINT NOT NULL
donated_at       TIMESTAMP NOT NULL

---

### access_grants table
campaign_id   UUID REFERENCES campaigns(id)
donor_wallet  VARCHAR(42) NOT NULL
granted_at    TIMESTAMP NOT NULL
PRIMARY KEY (campaign_id, donor_wallet)

---

### campaign_updates table
id            UUID PRIMARY KEY
campaign_id   UUID REFERENCES campaigns(id)
creator_wallet VARCHAR(42) NOT NULL
title         VARCHAR(100) NOT NULL
body          TEXT NOT NULL
update_type   VARCHAR(30) NOT NULL
ipfs_hashes   TEXT[]
update_hash   VARCHAR(66) NOT NULL
on_chain_tx   VARCHAR(66) NULLABLE
milestone_id  UUID NULLABLE
is_flagged    BOOLEAN DEFAULT FALSE
created_at    TIMESTAMP DEFAULT NOW()

---

### admin_actions table
id                 UUID PRIMARY KEY
action_type        VARCHAR(50) NOT NULL
campaign_id        UUID REFERENCES campaigns(id)
admin_wallet       VARCHAR(42) NOT NULL
notes              TEXT NULLABLE
safe_tx_hash       VARCHAR(66) NULLABLE
performed_at       TIMESTAMP DEFAULT NOW()

---

### email_queue table
id               UUID PRIMARY KEY
recipient_email  VARCHAR(255) NOT NULL
subject          VARCHAR(255) NOT NULL
template_name    VARCHAR(100) NOT NULL
template_data    JSONB NOT NULL
status           VARCHAR(20) DEFAULT 'pending'
attempts         INT DEFAULT 0
created_at       TIMESTAMP DEFAULT NOW()
sent_at          TIMESTAMP NULLABLE

---

## API Response Format
All responses:
{
  data: T | null,
  error: string | null,
  meta?: { total, page, limit }
}

---

## API Endpoints

### Public
GET  /api/campaigns
GET  /api/campaigns/:id
GET  /api/campaigns/:id/graph
GET  /api/campaigns/:id/updates
GET  /api/fundraiser/:wallet
GET  /api/health

---

### Donor
GET  /api/campaigns/:id/donors
GET  /api/campaigns/:id/access
GET  /api/donor/receipts

POST /api/campaigns/:id/prepare-donation
Body: { donorWallet, amountWei, txHash }

---

### Creator
POST /api/campaigns/:id/milestones/:idx/submit
POST /api/campaigns/:id/updates
POST /api/campaigns/:id/amendments/:idx

---

### Admin
GET  /api/admin/applications
POST /api/admin/campaigns/:id/approve
POST /api/admin/campaigns/:id/reject
POST /api/admin/campaigns/:id/freeze
POST /api/admin/campaigns/:id/milestones/:idx/approve
POST /api/admin/campaigns/:id/milestones/:idx/status
GET  /api/admin/safe/pending

---

### Auth
GET  /api/auth/nonce/:wallet
POST /api/auth/verify
POST /api/auth/logout

---

### WebSocket
GET /ws

---

## Build Order
1. config/config.go
2. db/postgres.go + migrations
3. models/
4. api/router.go + middleware
5. campaigns handler
6. blockchain client
7. watcher (critical)
8. IPFS upload
9. donation handler
10. websocket
11. email
12. admin
13. everything else