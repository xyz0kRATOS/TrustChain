-- 002_create_milestones.sql
-- Up migration: create the milestones table

CREATE TABLE IF NOT EXISTS milestones (
    id                UUID      PRIMARY KEY DEFAULT gen_random_uuid(),
    campaign_id       UUID      NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    sequence_index    INT       NOT NULL,
    description       TEXT      NOT NULL,
    amount_wei        NUMERIC(78) NOT NULL,
    deadline          TIMESTAMP NOT NULL,
    required_evidence TEXT      NOT NULL,
    status            VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    submitted_at      TIMESTAMP,
    approved_at       TIMESTAMP,
    released_at       TIMESTAMP,

    UNIQUE (campaign_id, sequence_index)
);

CREATE INDEX IF NOT EXISTS idx_milestones_campaign_id ON milestones(campaign_id);
CREATE INDEX IF NOT EXISTS idx_milestones_status      ON milestones(status);
