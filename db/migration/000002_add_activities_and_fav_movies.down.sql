ALTER TABLE IF EXISTS "activities" DROP CONSTRAINT IF EXISTS "activities_user_id_fkey";

ALTER TABLE IF EXISTS "fav_movies" DROP CONSTRAINT IF EXISTS "fav_movies_user_id_fkey";

ALTER TABLE IF EXISTS "sessions" DROP CONSTRAINT IF EXISTS "sessions_user_id_fkey";

DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS "activities";
DROP TABLE IF EXISTS "fav_movies";
DROP TABLE IF EXISTS "users";