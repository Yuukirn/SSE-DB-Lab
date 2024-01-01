USE bookstore_management_system;

DROP TABLE IF EXISTS user;

CREATE TABLE `user` (
    `id` INT AUTO_INCREMENT PRIMARY KEY NOT NULL COMMENT '用户ID',
    `name` VARCHAR(36) NOT NULL DEFAULT '' COMMENT '用户名',
    `password_hash` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码',
    `address` VARCHAR(255) DEFAULT NULL COMMENT '地址',
    `balance` DOUBLE NOT NULL DEFAULT 0.0 COMMENT '余额',
    `credit_rating` INT NOT NULL DEFAULT 1 COMMENT '信用等级',
    `email` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '邮箱',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_name` (`name`),
    KEY `idx_email` (`email`)
);
