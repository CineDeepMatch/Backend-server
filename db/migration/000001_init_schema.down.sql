ALTER TABLE IF EXISTS "sessions" DROP CONSTRAINT IF EXISTS "sessions_user_id_fkey";

DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS "users";