BEGIN;
--
-- Create model ContentType
--
CREATE TABLE "django_content_type" ("id" serial NOT NULL PRIMARY KEY, "name" varchar(100) NOT NULL, "app_label" varchar(100) NOT NULL, "model" varchar(100) NOT NULL);
--
-- Alter unique_together for contenttype (1 constraint(s))
--
ALTER TABLE "django_content_type" ADD CONSTRAINT "django_content_type_app_label_model_76bd3d3b_uniq" UNIQUE ("app_label", "model");
COMMIT;
BEGIN;
--
-- Change Meta options on contenttype
--
--
-- Alter field name on contenttype
--
ALTER TABLE "django_content_type" ALTER COLUMN "name" DROP NOT NULL;
--
-- MIGRATION NOW PERFORMS OPERATION THAT CANNOT BE WRITTEN AS SQL:
-- Raw Python operation
--
--
-- Remove field name from contenttype
--
ALTER TABLE "django_content_type" DROP COLUMN "name" CASCADE;
COMMIT;
