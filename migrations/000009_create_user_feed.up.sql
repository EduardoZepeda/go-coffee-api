BEGIN;
--
-- Create model Action
--
CREATE TABLE "feeds_action" ("id" bigserial NOT NULL PRIMARY KEY, "action" varchar(255) NOT NULL, "target_id" integer NULL CHECK ("target_id" >= 0), "created" timestamp with time zone NOT NULL, "target_ct_id" integer NULL, "user_id" bigint NOT NULL);
ALTER TABLE "feeds_action" ADD CONSTRAINT "feeds_action_target_ct_id_820f1b93_fk_django_content_type_id" FOREIGN KEY ("target_ct_id") REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE "feeds_action" ADD CONSTRAINT "feeds_action_user_id_05157ef6_fk_accounts_user_id" FOREIGN KEY ("user_id") REFERENCES "accounts_user" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "feeds_action_target_id_cb26698b" ON "feeds_action" ("target_id");
CREATE INDEX "feeds_action_created_66aa5327" ON "feeds_action" ("created");
CREATE INDEX "feeds_action_target_ct_id_820f1b93" ON "feeds_action" ("target_ct_id");
CREATE INDEX "feeds_action_user_id_05157ef6" ON "feeds_action" ("user_id");
COMMIT;
