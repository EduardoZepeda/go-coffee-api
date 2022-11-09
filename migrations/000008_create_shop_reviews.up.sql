BEGIN;
--
-- Create model Review
--
CREATE TABLE "reviews_review" ("id" bigserial NOT NULL PRIMARY KEY, "content" varchar(255) NOT NULL, "recommended" boolean NOT NULL, "created_date" timestamp with time zone NOT NULL, "modified_date" timestamp with time zone NOT NULL, "shop_id" bigint NOT NULL, "user_id" bigint NOT NULL);
--
-- Create constraint Only one review per user and shop on model review
--
ALTER TABLE "reviews_review" ADD CONSTRAINT "Only one review per user and shop" UNIQUE ("user_id", "shop_id");
--
-- Alter field shop on review
--
--
-- Alter field user on review
--
ALTER TABLE "reviews_review" ADD CONSTRAINT "reviews_review_shop_id_35a6a830_fk_shops_shop_id" FOREIGN KEY ("shop_id") REFERENCES "shops_shop" ("id") DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE "reviews_review" ADD CONSTRAINT "reviews_review_user_id_875caff2_fk_accounts_user_id" FOREIGN KEY ("user_id") REFERENCES "accounts_user" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "reviews_review_shop_id_35a6a830" ON "reviews_review" ("shop_id");
CREATE INDEX "reviews_review_user_id_875caff2" ON "reviews_review" ("user_id");
COMMIT;
