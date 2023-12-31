USE bookstore_management_system;

DROP TABLE IF EXISTS `supplier`;

CREATE TABLE `supplier` (
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    `name` varchar(45) NOT NULL,
    `address` varchar(45) NOT NULL,
    `email` varchar(45) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_name ON `supplier` (name);