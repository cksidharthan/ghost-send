-- name: CreateSecret :one
-- Creates a new secret with encrypted text and password protection
-- Args:
--   $1: secret text to encrypt
--   $2: password for encryption and access
--   $3: expiration timestamp
--   $4: number of allowed tries
INSERT INTO secrets (
    secret_text,
    password_hash,
    salt,
    expires_at,
    remaining_tries
) VALUES (
    pgp_sym_encrypt(sqlc.arg(secret_text), sqlc.arg(password))::text,
    crypt(sqlc.arg(password), gen_salt('bf'))::text,
    gen_salt('bf'),
    sqlc.arg(expires_at),
    sqlc.arg(remaining_tries)
) RETURNING *;

-- name: GetSecretByID :one
WITH secret_data AS (
    SELECT 
        s.*,
        (s.password_hash = crypt(sqlc.arg(password), s.password_hash)) as password_matches,
        CASE 
            WHEN remaining_tries <= 1 THEN TRUE
            ELSE FALSE
        END as should_delete
    FROM secrets s
    WHERE id = sqlc.arg(secret_id)::uuid 
    AND expires_at > CURRENT_TIMESTAMP
)
SELECT 
    id,
    CASE 
        WHEN password_matches THEN pgp_sym_decrypt(secret_text::bytea, sqlc.arg(password))::text 
        ELSE NULL 
    END::text as secret_text,
    password_hash,
    salt,
    expires_at,
    created_at,
    remaining_tries,
    last_viewed_at,
    password_matches,
    should_delete
FROM secret_data;

-- name: DeleteSecret :exec
-- Soft deletes a secret by UUID
-- Args: secret_id: secret UUID
DELETE FROM secrets
WHERE id = sqlc.arg(secret_id)::uuid;

-- name: DecrementTries :one
-- Decrements remaining tries and auto-deletes if no tries remain
-- Args: $1: secret UUID
UPDATE secrets
SET remaining_tries = remaining_tries - 1,
    last_viewed_at = CASE 
        WHEN remaining_tries <= 1 THEN CURRENT_TIMESTAMP 
        ELSE last_viewed_at 
    END
WHERE id = sqlc.arg(secret_id)::uuid
RETURNING *;

-- name: MarkSecretViewed :one
-- Marks a secret as viewed and deletes it (one-time view)
-- Args: $1: secret UUID
UPDATE secrets
SET last_viewed_at = CURRENT_TIMESTAMP,
    remaining_tries = 0
WHERE id = sqlc.arg(secret_id)::uuid
RETURNING *;

-- name: DeleteExpiredSecrets :exec
-- Deletes all expired secrets
DELETE FROM secrets
WHERE expires_at <= CURRENT_TIMESTAMP;

-- name: GetSecretStats :one
-- Returns statistics about secrets in the system:
-- - Number of active (non-deleted) secrets
-- - Number of viewed secrets
-- - Number of secrets that failed due to too many attempts
SELECT 
    COUNT(*) FILTER (WHERE remaining_tries > 0) as active_secrets,
    COUNT(*) FILTER (WHERE last_viewed_at IS NOT NULL) as viewed_secrets,
    COUNT(*) FILTER (WHERE remaining_tries = 0) as failed_attempts
FROM secrets;

-- name: CheckSecretStatus :one
-- Checks if a secret exists and is still accessible
-- Args: $1: secret UUID
SELECT EXISTS (
    SELECT 1 FROM secrets 
    WHERE id = sqlc.arg(secret_id)::uuid 
    AND expires_at > CURRENT_TIMESTAMP
    AND remaining_tries > 0
) as exists;
