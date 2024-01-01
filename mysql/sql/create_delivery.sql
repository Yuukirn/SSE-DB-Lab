USE bookstore_management_system;

DROP TABLE IF EXISTS `delivery`;

CREATE TABLE `delivery` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键 id',
  `order_id` int(11) NOT NULL COMMENT '订单 id',
  `user_id` int(11) NOT NULL COMMENT '用户 id',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  FOREIGN KEY (`order_id`) REFERENCES `order` (`id`)
);