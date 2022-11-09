BEGIN;
--
-- Create model CoffeeBag
--
DROP TABLE "shops_coffeebag_coffee_shop" CASCADE;
DROP TABLE "shops_coffeebag" CASCADE;
--
-- Create model Shop
--
DROP TABLE "shops_shop_likes" CASCADE;
DROP TABLE "shops_shop" CASCADE;
COMMIT;
