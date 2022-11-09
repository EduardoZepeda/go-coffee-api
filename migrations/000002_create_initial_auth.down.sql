BEGIN;
--
-- Create model User
--
--
-- Create model Group
--
DROP TABLE "auth_group_permissions" CASCADE;
DROP TABLE "auth_group" CASCADE;
--
-- Create model Permission
--
DROP TABLE "auth_permission" CASCADE;
COMMIT;