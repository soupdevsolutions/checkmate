CREATE TABLE IF NOT EXISTS healthchecks(
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "target_id" UUID NOT NULL REFERENCES "targets"(id),
    "status" TEXT NOT NULL,
    "timestamp" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);