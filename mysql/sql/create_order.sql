USE bookstore_management_system;

DROP TABLE IF EXISTS `order`;

CREATE TABLE `order` (
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '订单id',
    `order_date` date NOT NULL COMMENT '订单日期',
    `user_id` int NOT NULL COMMENT '用户id',
    `book_ids` JSON NOT NULL COMMENT '书籍id',
    `order_amounts` JSON NOT NULL COMMENT '书籍数量',
    `order_money` double NOT NULL COMMENT '订单金额',
    `delivery_address` varchar(255) NOT NULL COMMENT '收货地址',
    `delivery_status` varchar(255) NOT NULL COMMENT '发货状态',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_user_id` (`user_id`),
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
);
