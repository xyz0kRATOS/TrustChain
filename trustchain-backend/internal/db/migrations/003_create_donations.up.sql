-- 003_create_donations.sql
-- Up migration: create the donations table

CREATE TABLE IF NOT EXISTS donations (
    id                UUID      PRIMARY KEY DEFAULT gen_random_uuid(),
    campaign_id       UUID      NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    donor_wallet      VARCHAR(42)  NOT NULL,
    amount_wei        NUMERIC(78)  NOT NULL,
    tx_hash           VARCHAR(66)  NOT NULL,
    nft_token_id      INT,
    ipfs_metadata_cid TEXT,
    block_number      BIGINT       NOT NULL,
    donated_at        TIMESTAMP    NOT NULL,

    UNIQUE (tx_hash)
);

CREATE INDEX IF NOT EXISTS idx_donations_campaign_id  ON donations(campaign_id);
CREATE INDEX IF NOT EXISTS idx_donations_donor_wallet ON donations(donor_wallet);
CREATE INDEX IF NOT EXISTS idx_donations_tx_hash      ON donations(tx_hash);
