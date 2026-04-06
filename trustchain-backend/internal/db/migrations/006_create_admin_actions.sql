-- 006_create_admin_actions.sql
-- Up migration: create the admin_actions table

CREATE TABLE IF NOT EXISTS admin_actions (
    id           UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    action_type  VARCHAR(50) NOT NULL,
    campaign_id  UUID        REFERENCES campaigns(id) ON DELETE SET NULL,
    admin_wallet VARCHAR(42) NOT NULL,
    notes        TEXT,
    safe_tx_hash VARCHAR(66),
    performed_at TIMESTAMP   NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_admin_actions_campaign_id ON admin_actions(campaign_id);
CREATE INDEX IF NOT EXISTS idx_admin_actions_admin_wallet ON admin_actions(admin_wallet);
