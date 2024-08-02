
-- +migrate Up
ALTER TABLE `stories` ADD `category_id` varchar(50) NOT NULL;

-- +migrate Down
ALTER TABLE `stories` DROP COLUMN `category_id`;
