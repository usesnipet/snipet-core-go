-- Create "knowledge_bases" table
CREATE TABLE "knowledge_bases" (
  "id" uuid NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_knowledge_bases_deleted_at" to table: "knowledge_bases"
CREATE INDEX "idx_knowledge_bases_deleted_at" ON "knowledge_bases" ("deleted_at");
-- Create index "idx_knowledge_bases_name" to table: "knowledge_bases"
CREATE UNIQUE INDEX "idx_knowledge_bases_name" ON "knowledge_bases" ("name");
-- Create "knowledge_sources" table
CREATE TABLE "knowledge_sources" (
  "id" uuid NOT NULL,
  "name" character varying(255) NOT NULL,
  "provider" character varying(50) NOT NULL,
  "provider_type" character varying(20) NOT NULL,
  "config" jsonb NOT NULL,
  "use_rag" boolean NULL DEFAULT true,
  "rag_strategy" character varying(20) NULL,
  "rag_config" jsonb NULL,
  "status" character varying(20) NULL DEFAULT 'active',
  "last_sync_at" timestamptz NULL,
  "last_error" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_knowledge_sources_deleted_at" to table: "knowledge_sources"
CREATE INDEX "idx_knowledge_sources_deleted_at" ON "knowledge_sources" ("deleted_at");
-- Create "knowledge_base_sources" table
CREATE TABLE "knowledge_base_sources" (
  "knowledge_base_id" uuid NOT NULL,
  "knowledge_source_id" uuid NOT NULL,
  PRIMARY KEY ("knowledge_base_id", "knowledge_source_id"),
  CONSTRAINT "fk_knowledge_base_sources_knowledge_base" FOREIGN KEY ("knowledge_base_id") REFERENCES "knowledge_bases" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_knowledge_base_sources_knowledge_source" FOREIGN KEY ("knowledge_source_id") REFERENCES "knowledge_sources" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
