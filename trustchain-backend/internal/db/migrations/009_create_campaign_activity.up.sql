CREATE TABLE IF NOT EXISTS campaign_activity (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type          VARCHAR(50) NOT NULL,
    campaign_id   UUID REFERENCES campaigns(id),
    campaign_name VARCHAR(255),
    wallet        VARCHAR(42),
    amount_wei    NUMERIC(78),
    tx_hash       VARCHAR(66),
    block_number  BIGINT,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_campaign_activity_created_at ON campaign_activity(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_campaign_activity_campaign_id ON campaign_activity(campaign_id);
