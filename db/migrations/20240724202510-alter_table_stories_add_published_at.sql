
-- +migrate Up
ALTER TABLE `stories` ADD `published_at` TIMESTAMP NULL;
ALTER TABLE `stories` ADD `created_at` TIMESTAMP NOT NULL DEFAULT NOW();

-- +migrate Down
ALTER TABLE `stories` DROP COLUMN `published_at`; 
ALTER TABLE `stories` DROP COLUMN `created_at`; 
