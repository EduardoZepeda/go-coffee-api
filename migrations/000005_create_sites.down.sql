BEGIN;
--
-- Alter field domain on site
--
ALTER TABLE "django_site" DROP CONSTRAINT "django_site_domain_a2e37b91_uniq";
DROP INDEX IF EXISTS "django_site_domain_a2e37b91_like";
COMMIT;
BEGIN;
--
-- Create model Site
--
DROP TABLE "django_site" CASCADE;
COMMIT;
