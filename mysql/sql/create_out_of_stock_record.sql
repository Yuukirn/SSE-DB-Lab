USE bookstore_management_system;

DROP TABLE IF EXISTS `out_of_stack_record`;

CREATE TABLE `out_of_stock_record` (
    `id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键 id',
    `book_id` INT NOT NULL DEFAULT 0 COMMENT '书籍 id',
    `number` INT NOT NULL DEFAULT 0 COMMENT '缺货数量',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `book_id` (`book_id`),
    FOREIGN KEY (`book_id`) REFERENCES `book`(`id`)
);
