BEGIN;
--
-- Create constraint unique-email on model user
--
ALTER TABLE "accounts_user" DROP CONSTRAINT "unique-email";
--
-- Create constraint email-user on model user
--
ALTER TABLE "accounts_user" DROP CONSTRAINT "email-user";
--
-- Change Meta options on user
--
--
-- Create constraint userTo-userFrom on model contact
--
ALTER TABLE "accounts_contact" DROP CONSTRAINT "userTo-userFrom";
--
-- Add field bio to user
--
ALTER TABLE "accounts_user" DROP COLUMN "bio" CASCADE;
--
-- Add field profile_picture to user
--
ALTER TABLE "accounts_user" DROP COLUMN "profile_picture" CASCADE;
--
-- Add field user_permissions to user
--
DROP TABLE "accounts_user_user_permissions" CASCADE;
--
-- Add field groups to user
--
DROP TABLE "accounts_user_groups" CASCADE;
--
-- Add field following to user
--
--
-- Create model Contact
--
DROP TABLE "accounts_contact" CASCADE;
--
-- Create model User
--
DROP TABLE "accounts_user" CASCADE;
COMMIT;
