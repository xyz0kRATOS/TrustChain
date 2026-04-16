ALTER TABLE milestones
  DROP COLUMN IF EXISTS amount_usd,
  DROP COLUMN IF EXISTS name;

ALTER TABLE campaigns
  DROP COLUMN IF EXISTS document_file_names,
  DROP COLUMN IF EXISTS creator_bio,
  DROP COLUMN IF EXISTS creator_country,
  DROP COLUMN IF EXISTS creator_org,
  DROP COLUMN IF EXISTS creator_email,
  DROP COLUMN IF EXISTS creator_name,
  DROP COLUMN IF EXISTS goal_amount_usd,
  DROP COLUMN IF EXISTS category;
