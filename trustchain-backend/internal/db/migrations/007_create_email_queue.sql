-- 007_create_email_queue.sql
-- Up migration: create the email_queue table

CREATE TABLE IF NOT EXISTS email_queue (
    id              UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    recipient_email VARCHAR(255) NOT NULL,
    subject         VARCHAR(255) NOT NULL,
    template_name   VARCHAR(100) NOT NULL,
    template_data   JSONB        NOT NULL DEFAULT '{}'::jsonb,
    status          VARCHAR(20)  NOT NULL DEFAULT 'pending',
    attempts        INT          NOT NULL DEFAULT 0,
    created_at      TIMESTAMP    NOT NULL DEFAULT NOW(),
    sent_at         TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_email_queue_status ON email_queue(status);
