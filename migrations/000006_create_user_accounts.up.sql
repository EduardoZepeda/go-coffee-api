BEGIN;
--
-- Create model User
--
CREATE TABLE "accounts_user" ("id" bigserial NOT NULL PRIMARY KEY, "password" varchar(128) NOT NULL, "last_login" timestamp with time zone NULL, "is_superuser" boolean NOT NULL, "username" varchar(150) NOT NULL UNIQUE, "first_name" varchar(150) NOT NULL, "last_name" varchar(150) NOT NULL, "email" varchar(254) NOT NULL, "is_staff" boolean NOT NULL, "is_active" boolean NOT NULL, "date_joined" timestamp with time zone NOT NULL);
--
-- Create model Contact
--
CREATE TABLE "accounts_contact" ("id" bigserial NOT NULL PRIMARY KEY, "created" timestamp with time zone NOT NULL, "user_from_id" bigint NOT NULL, "user_to_id" bigint NOT NULL);
--
-- Add field following to user
--
--
-- Add field groups to user
--
CREATE TABLE "accounts_user_groups" ("id" bigserial NOT NULL PRIMARY KEY, "user_id" bigint NOT NULL, "group_id" integer NOT NULL);
--
-- Add field user_permissions to user
--
CREATE TABLE "accounts_user_user_permissions" ("id" bigserial NOT NULL PRIMARY KEY, "user_id" bigint NOT NULL, "permission_id" integer NOT NULL);
--
-- Add field profile_picture to user
--
ALTER TABLE "accounts_user" ADD COLUMN "profile_picture" varchar(100) NULL;
--
-- Add field bio to user
--
ALTER TABLE "accounts_user" ADD COLUMN "bio" varchar(250) NULL;
--
-- Create constraint userTo-userFrom on model contact
--
ALTER TABLE "accounts_contact" ADD CONSTRAINT "userTo-userFrom" UNIQUE ("user_to_id", "user_from_id");
--
-- Change Meta options on user
--
--
-- Create constraint email-user on model user
--
ALTER TABLE "accounts_user" ADD CONSTRAINT "email-user" UNIQUE ("username", "email");
--
-- Create constraint unique-email on model user
--
ALTER TABLE "accounts_user" ADD CONSTRAINT "unique-email" UNIQUE ("email");
CREATE INDEX "accounts_user_username_6088629e_like" ON "accounts_user" ("username" varchar_pattern_ops);
ALTER TABLE "accounts_contact" ADD CONSTRAINT "accounts_contact_user_from_id_d88fc381_fk_accounts_user_id" FOREIGN KEY ("user_from_id") REFERENCES "accounts_user" ("id") DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE "accounts_contact" ADD CONSTRAINT "accounts_contact_user_to_id_16d11cce_fk_accounts_user_id" FOREIGN KEY ("user_to_id") REFERENCES "accounts_user" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "accounts_contact_created_5be07012" ON "accounts_contact" ("created");
CREATE INDEX "accounts_contact_user_from_id_d88fc381" ON "accounts_contact" ("user_from_id");
CREATE INDEX "accounts_contact_user_to_id_16d11cce" ON "accounts_contact" ("user_to_id");
ALTER TABLE "accounts_user_groups" ADD CONSTRAINT "accounts_user_groups_user_id_group_id_59c0b32f_uniq" UNIQUE ("user_id", "group_id");
ALTER TABLE "accounts_user_groups" ADD CONSTRAINT "accounts_user_groups_user_id_52b62117_fk_accounts_user_id" FOREIGN KEY ("user_id") REFERENCES "accounts_user" ("id") DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE "accounts_user_groups" ADD CONSTRAINT "accounts_user_groups_group_id_bd11a704_fk_auth_group_id" FOREIGN KEY ("group_id") REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "accounts_user_groups_user_id_52b62117" ON "accounts_user_groups" ("user_id");
CREATE INDEX "accounts_user_groups_group_id_bd11a704" ON "accounts_user_groups" ("group_id");
ALTER TABLE "accounts_user_user_permissions" ADD CONSTRAINT "accounts_user_user_permi_user_id_permission_id_2ab516c2_uniq" UNIQUE ("user_id", "permission_id");
ALTER TABLE "accounts_user_user_permissions" ADD CONSTRAINT "accounts_user_user_p_user_id_e4f0a161_fk_accounts_" FOREIGN KEY ("user_id") REFERENCES "accounts_user" ("id") DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE "accounts_user_user_permissions" ADD CONSTRAINT "accounts_user_user_p_permission_id_113bb443_fk_auth_perm" FOREIGN KEY ("permission_id") REFERENCES "auth_permission" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "accounts_user_user_permissions_user_id_e4f0a161" ON "accounts_user_user_permissions" ("user_id");
CREATE INDEX "accounts_user_user_permissions_permission_id_113bb443" ON "accounts_user_user_permissions" ("permission_id");
COMMIT;
