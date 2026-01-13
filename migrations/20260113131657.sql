-- Modify "transactions" table
ALTER TABLE "public"."transactions" ADD COLUMN "category_id" bigint NULL, ADD CONSTRAINT "fk_transactions_category" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Create index "idx_transactions_category_id" to table: "transactions"
CREATE INDEX "idx_transactions_category_id" ON "public"."transactions" ("category_id");
