
-- +migrate Up
CREATE TABLE `profiles`(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    `bio` text,
    `image` varchar(255),
    
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);
-- +migrate Down
DROP TABLE IF EXISTS profiles;
