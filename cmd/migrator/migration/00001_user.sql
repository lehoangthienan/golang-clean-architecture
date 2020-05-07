-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."users"(
	"id" uuid NOT NULL,
	"created_at" timestamptz DEFAULT now(),
	"deleted_at" timestamptz,
	"updated_at" timestamptz,
	"name" text NOT NULL,
	"user_name" text NOT NULL,
	"pass_word" text NOT NULL,
	"role" text NOT NULL
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."users";