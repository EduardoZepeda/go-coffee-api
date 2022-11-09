BEGIN;
--
-- Alter field user on review
--
--
-- Alter field shop on review
--
--
-- Create constraint Only one review per user and shop on model review
--
ALTER TABLE "reviews_review" DROP CONSTRAINT "Only one review per user and shop";
--
-- Create model Review
--
DROP TABLE "reviews_review" CASCADE;
COMMIT;
