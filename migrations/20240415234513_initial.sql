-- Create "todos" table
CREATE TABLE "todos" ("id" uuid NOT NULL, "title" character varying(50) NOT NULL, "description" character varying(1000) NULL, "completed" boolean NOT NULL, "created_at" timestamp NOT NULL, "updated_at" timestamp NOT NULL, PRIMARY KEY ("id"));
