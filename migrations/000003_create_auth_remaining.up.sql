BEGIN;
--
-- Alter field name on permission
--
ALTER TABLE "auth_permission" ALTER COLUMN "name" TYPE varchar(255);
COMMIT;
BEGIN;
--
-- Alter field email on user
--
COMMIT;
BEGIN;
--
-- Alter field username on user
--
COMMIT;
BEGIN;
--
-- Alter field last_login on user
--
COMMIT;
BEGIN;
--
-- Alter field username on user
--
COMMIT;
BEGIN;
--
-- Alter field username on user
--
COMMIT;
BEGIN;
--
-- Alter field last_name on user
--
COMMIT;
BEGIN;
--
-- Alter field name on group
--
ALTER TABLE "auth_group" ALTER COLUMN "name" TYPE varchar(150);
COMMIT;
BEGIN;
--
-- MIGRATION NOW PERFORMS OPERATION THAT CANNOT BE WRITTEN AS SQL:
-- Raw Python operation
--
COMMIT;
BEGIN;
--
-- Alter field first_name on user
--
COMMIT;
