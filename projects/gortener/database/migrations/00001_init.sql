-- +goose Up
-- +goose StatementBegin
CREATE TABLE "urls" (
  "id" integer PRIMARY KEY,
  "url" varchar,
  "short_url" varchar UNIQUE,
  "hash" varchar,
  "created_at" timestamp DEFAULT (now()),
  "user_id" integer
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar,
  "name" varchar,
  "email" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

COMMENT ON TABLE "urls" IS 'Store the urls created by a user';

COMMENT ON TABLE "users" IS 'Users that perform any action in the application';

ALTER TABLE "urls" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS urls;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
