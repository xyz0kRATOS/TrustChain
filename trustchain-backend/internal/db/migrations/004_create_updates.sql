-- 004_create_updates.sql
-- Up migration: create the campaign_updates table

CREATE TABLE IF NOT EXISTS campaign_updates (
    id             UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    campaign_id    UUID        NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    creator_wallet VARCHAR(42) NOT NULL,
    title          VARCHAR(100) NOT NULL,
    body           TEXT         NOT NULL,
    update_type    VARCHAR(30)  NOT NULL,
    ipfs_hashes    TEXT[]       NOT NULL DEFAULT '{}',
    update_hash    VARCHAR(66)  NOT NULL,
    on_chain_tx    VARCHAR(66),
    milestone_id   UUID         REFERENCES milestones(id),
    is_flagged     BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at     TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_campaign_updates_campaign_id ON campaign_updates(campaign_id);
CREATE INDEX IF NOT EXISTS idx_campaign_updates_creator     ON campaign_updates(creator_wallet);
