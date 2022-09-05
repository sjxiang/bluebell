CREATE DATABASE IF NOT EXISTS `bulebell` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `bluebell`;

DROP TABLE IF EXISTS `user`;


-- 用户表结构设计

CREATE TABLE `user` (
    `id`          bigint(20)  NOT NULL AUTO_INCREMENT,
    `user_id`     varchar(64) NOT NULL,
    `username`    varchar(64) NOT NULL,
    `password`    varchar(64),
    `email`       varchar(64),
    `gender`      tinyint(4) DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


describe user;


