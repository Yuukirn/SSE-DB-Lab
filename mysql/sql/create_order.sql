USE bookstore_management_system;

DROP TABLE IF EXISTS `order`;

CREATE TABLE `order` (
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    `order_date` date NOT NULL,
    `user_id` int NOT NULL,
    `book_ids` JSON NOT NULL,
    `order_amounts` JSON NOT NULL,
    `order_money` double NOT NULL,
    `delivery_address` varchar(255) NOT NULL,
    `delivery_status` varchar(255) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_user_id ON `order` (user_id);