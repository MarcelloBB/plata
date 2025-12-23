-- Modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "created_at", DROP COLUMN "updated_at", DROP COLUMN "deleted_at";
-- Create "budgets" table
CREATE TABLE "public"."budgets" (
  "id" bigserial NOT NULL,
  "user_id" bigint NOT NULL,
  "amount" numeric NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_budgets_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_budgets_user_id" to table: "budgets"
CREATE INDEX "idx_budgets_user_id" ON "public"."budgets" ("user_id");
-- Create "categories" table
CREATE TABLE "public"."categories" (
  "id" bigserial NOT NULL,
  "name" character varying(50) NOT NULL,
  "user_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_categories_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_categories_user_id" to table: "categories"
CREATE INDEX "idx_categories_user_id" ON "public"."categories" ("user_id");
-- Create "transactions" table
CREATE TABLE "public"."transactions" (
  "id" bigserial NOT NULL,
  "user_id" bigint NOT NULL,
  "amount" numeric NULL,
  "description" character varying(75) NULL,
  "date" timestamptz NULL,
  "type" character varying(10) NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_transactions_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_transactions_user_id" to table: "transactions"
CREATE INDEX "idx_transactions_user_id" ON "public"."transactions" ("user_id");
