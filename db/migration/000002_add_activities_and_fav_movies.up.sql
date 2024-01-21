ALTER TABLE IF EXISTS "sessions" DROP CONSTRAINT IF EXISTS "sessions_user_id_fkey";

DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS "users";

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "create_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "activities" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "view_page" varchar NOT NULL,
  "duration" int NOT NULL,
  "page_visited_at" timestamptz NOT NULL
);

CREATE TABLE "fav_movies" (
  "user_id" uuid NOT NULL PRIMARY KEY,
  "movies" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "sessions" ("user_id", "create_at");

CREATE INDEX ON "activities" ("user_id");

CREATE INDEX ON "activities" ("user_id", "page_visited_at");

CREATE INDEX ON "fav_movies" ("user_id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "activities" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "fav_movies" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
