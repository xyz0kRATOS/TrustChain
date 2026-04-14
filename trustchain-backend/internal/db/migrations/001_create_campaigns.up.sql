-- 001_create_campaigns.sql
-- Up migration: create the campaigns table

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS campaigns (
    id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    contract_address   VARCHAR(42),
    creator_wallet     VARCHAR(42)  NOT NULL,
    name               VARCHAR(255) NOT NULL,
    description        TEXT         NOT NULL,
    goal_amount_wei    NUMERIC(78)  NOT NULL,
    status             VARCHAR(20)  NOT NULL DEFAULT 'PENDING',
    document_hash      VARCHAR(66),
    image_url          TEXT,
    ipfs_evidence_hash TEXT,
    created_at         TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_campaigns_creator_wallet ON campaigns(creator_wallet);
CREATE INDEX IF NOT EXISTS idx_campaigns_status          ON campaigns(status);
CREATE INDEX IF NOT EXISTS idx_campaigns_contract_address ON campaigns(contract_address);
