-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."groups"(
	"id" uuid NOT NULL,
	"created_at" timestamptz DEFAULT now(),
	"updated_at" timestamptz,
	"deleted_at" timestamptz,
    "name" text NOT NULL,
	"creator_id" uuid NOT NULL,
	CONSTRAINT "groups_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."groups";