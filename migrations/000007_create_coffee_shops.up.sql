BEGIN;
--
-- Create model Shop
--
CREATE TABLE "shops_shop" ("id" bigserial NOT NULL PRIMARY KEY, "name" varchar(100) NOT NULL, "location" geometry(POINT,4326) NOT NULL, "address" varchar(100) NOT NULL, "city" varchar(50) NOT NULL, "roaster" boolean NOT NULL, "rating" numeric(2, 1) NOT NULL, "created_date" timestamp with time zone NOT NULL, "modified_date" timestamp with time zone NOT NULL, "content" text NULL);
CREATE TABLE "shops_shop_likes" ("id" bigserial NOT NULL PRIMARY KEY, "shop_id" bigint NOT NULL, "user_id" bigint NOT NULL);
--
-- Create model CoffeeBag
--
CREATE TABLE "shops_coffeebag" ("id" bigserial NOT NULL PRIMARY KEY, "brand" varchar(200) NOT NULL, "species" varchar(2) NOT NULL, "origin" varchar(2) NULL);
CREATE TABLE "shops_coffeebag_coffee_shop" ("id" bigserial NOT NULL PRIMARY KEY, "coffeebag_id" bigint NOT NULL, "shop_id" bigint NOT NULL);
CREATE INDEX "shops_shop_location_523240f6_id" ON "shops_shop" USING GIST ("location");
ALTER TABLE "shops_shop_likes" ADD CONSTRAINT "shops_shop_likes_shop_id_user_id_09e87394_uniq" UNIQUE ("shop_id", "user_id");
ALTER TABLE "shops_shop_likes" ADD CONSTRAINT "shops_shop_likes_shop_id_c364c364_fk_shops_shop_id" FOREIGN KEY ("shop_id") REFERENCES "shops_shop" ("id") DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE "shops_shop_likes" ADD CONSTRAINT "shops_shop_likes_user_id_6f0a5822_fk_accounts_user_id" FOREIGN KEY ("user_id") REFERENCES "accounts_user" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "shops_shop_likes_shop_id_c364c364" ON "shops_shop_likes" ("shop_id");
CREATE INDEX "shops_shop_likes_user_id_6f0a5822" ON "shops_shop_likes" ("user_id");
ALTER TABLE "shops_coffeebag_coffee_shop" ADD CONSTRAINT "shops_coffeebag_coffee_shop_coffeebag_id_shop_id_2d92af17_uniq" UNIQUE ("coffeebag_id", "shop_id");
ALTER TABLE "shops_coffeebag_coffee_shop" ADD CONSTRAINT "shops_coffeebag_coff_coffeebag_id_f7e31f2a_fk_shops_cof" FOREIGN KEY ("coffeebag_id") REFERENCES "shops_coffeebag" ("id") DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE "shops_coffeebag_coffee_shop" ADD CONSTRAINT "shops_coffeebag_coffee_shop_shop_id_3f6bc214_fk_shops_shop_id" FOREIGN KEY ("shop_id") REFERENCES "shops_shop" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "shops_coffeebag_coffee_shop_coffeebag_id_f7e31f2a" ON "shops_coffeebag_coffee_shop" ("coffeebag_id");
CREATE INDEX "shops_coffeebag_coffee_shop_shop_id_3f6bc214" ON "shops_coffeebag_coffee_shop" ("shop_id");
COMMIT;
