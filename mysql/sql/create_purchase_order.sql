USE bookstore_management_system;

DROP TABLE IF EXISTS `purchase_order`;

CREATE TABLE `purchase_order` (
    `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    `out_of_stock_record_id` int(11) NOT NULL,
    `number` int(11) NOT NULL,
    `status` varchar(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY `idx_out_of_stock_record_id` (`out_of_stock_record_id`),
    CONSTRAINT `purchase_order_ibfk_1` FOREIGN KEY (`out_of_stock_record_id`) REFERENCES `out_of_stock_record` (`id`)
);