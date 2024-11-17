-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."patient";

CREATE TABLE "public"."patient" (
    "id" UUID NOT NULL DEFAULT gen_random_uuid(),
    -- Assuming Base has an id field
    "first_name" VARCHAR,
    -- Corresponds to *string (nullable)
    "last_name" VARCHAR,
    -- Corresponds to *string (nullable)
    "age" INT8 NOT NULL,
    -- Corresponds to int64
    "email" VARCHAR,
    -- Corresponds to *string (nullable)
    "phone" VARCHAR NOT NULL,
    -- Corresponds to string
    "disease" VARCHAR NOT NULL,
    -- Corresponds to string
    "address" JSONB,
    -- Corresponds to Address (assumed to be a JSON structure)
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    -- Assuming BaseAudit has created_at
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    -- Assuming BaseAudit has updated_at
    PRIMARY KEY ("id")
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."patient";

-- +goose StatementEnd