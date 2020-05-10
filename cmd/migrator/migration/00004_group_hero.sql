-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."group_heros"(
	"id" uuid NOT NULL,
	"created_at" timestamptz DEFAULT now(),
	"updated_at" timestamptz,
	"deleted_at" timestamptz,
	"group_id" uuid references groups,
	"hero_id" uuid references heros,
	CONSTRAINT "group_heros_pkey" PRIMARY KEY ("id"),
	UNIQUE("group_id", "hero_id")
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."group_heros";