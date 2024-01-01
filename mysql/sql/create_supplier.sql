USE bookstore_management_system;

DROP TABLE IF EXISTS `supplier`;

CREATE TABLE `supplier` (
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '供应商 id',
    `name` varchar(45) NOT NULL COMMENT '供应商名称',
    `address` varchar(45) NOT NULL COMMENT '供应商地址',
    `email` varchar(45) NOT NULL COMMENT '供应商邮箱',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_name` (`name`)
);
