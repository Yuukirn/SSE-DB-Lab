USE bookstore_management_system;

DROP TABLE IF EXISTS `purchase_order`;

CREATE TABLE `purchase_order` (
    `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '采购单id',
    `out_of_stock_record_id` int(11) NOT NULL COMMENT '缺货记录id',
    `number` int(11) NOT NULL COMMENT '采购数量',
    `status` varchar(255) NOT NULL COMMENT '采购单状态',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_out_of_stock_record_id` (`out_of_stock_record_id`),
    CONSTRAINT `purchase_order_ibfk_1` FOREIGN KEY (`out_of_stock_record_id`) REFERENCES `out_of_stock_record` (`id`)
);