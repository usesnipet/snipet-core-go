-- Create "knowledges" table
CREATE TABLE "knowledges" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" character varying(255) NULL,
  "description" character varying(255) NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_knowledges_deleted_at" to table: "knowledges"
CREATE INDEX "idx_knowledges_deleted_at" ON "knowledges" ("deleted_at");
