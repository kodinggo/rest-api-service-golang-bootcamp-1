CREATE TABLE `stories`(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `title` varchar(50) NOT NULL,
    `content` text NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `users`(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `username` varchar(50) NOT NULL,
    `password`  varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
);