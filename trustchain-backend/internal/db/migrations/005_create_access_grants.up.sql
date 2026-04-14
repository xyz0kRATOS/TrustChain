-- 005_create_access_grants.sql
-- Up migration: create the access_grants table

CREATE TABLE IF NOT EXISTS access_grants (
    campaign_id  UUID        NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    donor_wallet VARCHAR(42) NOT NULL,
    granted_at   TIMESTAMP   NOT NULL DEFAULT NOW(),

    PRIMARY KEY (campaign_id, donor_wallet)
);

CREATE INDEX IF NOT EXISTS idx_access_grants_donor_wallet ON access_grants(donor_wallet);
