BEGIN;
--
-- Create model Site
--
CREATE TABLE "django_site" ("id" serial NOT NULL PRIMARY KEY, "domain" varchar(100) NOT NULL, "name" varchar(50) NOT NULL);
COMMIT;
BEGIN;
--
-- Alter field domain on site
--
ALTER TABLE "django_site" ADD CONSTRAINT "django_site_domain_a2e37b91_uniq" UNIQUE ("domain");
CREATE INDEX "django_site_domain_a2e37b91_like" ON "django_site" ("domain" varchar_pattern_ops);
COMMIT;
