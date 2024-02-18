-- +goose Up
-- +goose StatementBegin
CREATE TABLE "urls" (
  "id" BIGSERIAL PRIMARY KEY, -- BIGSERIAL is pseudo-type more info https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-SERIAL
  "url" VARCHAR NOT NULL,
  "short_url" VARCHAR NOT NULL UNIQUE,
  "hash" VARCHAR NOT NULL,
  "created_at" TIMESTAMP DEFAULT (now()),
  "user_id" INTEGER
);

CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY, -- BIGSERIAL is pseudo-type more info https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-SERIAL
  "username" VARCHAR NOT NULL UNIQUE,
  "name" VARCHAR,
  "email" VARCHAR,
  "created_at" TIMESTAMP DEFAULT (now()),
  "updated_at" TIMESTAMP DEFAULT (now())
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
