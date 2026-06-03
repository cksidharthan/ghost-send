CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS secrets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    secret_text TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    salt TEXT NOT NULL DEFAULT gen_salt('bf'),
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    remaining_tries INTEGER NOT NULL DEFAULT 3,
    last_viewed_at TIMESTAMP WITH TIME ZONE
);

-- Add indexes for better query performance
-- Expires_at index helps with cleanup queries
CREATE INDEX idx_secrets_expires_at ON secrets(expires_at);
-- Created_at index helps with analytics and maintenance
CREATE INDEX idx_secrets_created_at ON secrets(created_at);