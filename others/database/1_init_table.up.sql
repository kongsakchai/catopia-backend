SET NAMES utf8mb4;

SET time_zone = '+07:00';

-- User
CREATE TABLE `users` (
    `id` int NOT NULL AUTO_INCREMENT,
    `username` varchar(32) NOT NULL,
    `password` varchar(128) NOT NULL,
    `email` varchar(64) NOT NULL,
    `salt` varchar(16) NOT NULL,
    `gender` varchar(8) NOT NULL,
    `profile` varchar(64) DEFAULT NULL,
    `date` date NOT NULL,
    `group_id` int DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `sessions` (
    `id` varchar(64) NOT NULL,
    `user_id` int NOT NULL,
    `token` varchar(256) NOT NULL,
    `expired_at` timestamp NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);

-- Cat
CREATE TABLE `cat` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(32) NOT NULL,
    `weight` float NOT NULL,
    `gender` varchar(8) NOT NULL,
    `profile` varchar(64) DEFAULT NULL,
    `date` date NOT NULL,
    `breeding` varchar(32) NOT NULL,
    `aggression` int NOT NULL DEFAULT '0',
    `shyness` int NOT NULL DEFAULT '0',
    `extraversion` int NOT NULL DEFAULT '0',
    `user_id` int NOT NULL,
    `group_id` int DEFAULT NULL,
    `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);

-- Breed
CREATE TABLE `breeding` (
    `id` int NOT NULL AUTO_INCREMENT,
    `breed` varchar(64) NOT NULL,
    `group_name` varchar(4) NOT NULL,
    PRIMARY KEY (`id`)
);

-- Group
CREATE TABLE `cat_group` (
    `id` int NOT NULL AUTO_INCREMENT,
    `group_name` varchar(4) NOT NULL,
    `count` int NOT NULL DEFAULT 0,
    `group` int NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

-- Teatment
CREATE TABLE `treatment_type` (
    `id` int NOT NULL AUTO_INCREMENT,
    `treatment_type` varchar(64) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `treatment` (
    `id` int NOT NULL AUTO_INCREMENT,
    `cat_id` int NOT NULL,
    `treatment_type_id` int DEFAULT NULL,
    `date` date NOT NULL,
    `location` varchar(64) DEFAULT NULL,
    `vet` varchar(64) DEFAULT NULL,
    `detail` varchar(64) DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`cat_id`) REFERENCES `cat` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`treatment_type_id`) REFERENCES `treatment_type` (`id`) ON DELETE
    SET NULL
);