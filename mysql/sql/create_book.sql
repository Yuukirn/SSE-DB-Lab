USE bookstore_management_system;

DROP TABLE IF EXISTS `book`;

CREATE TABLE `book` (
    `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT '书籍ID',
    `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '书名',
    `authors` JSON NOT NULL COMMENT '作者',
    `publisher` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '出版社',
    `price` DOUBLE NOT NULL DEFAULT 0.0 COMMENT '价格',
    `keywords` JSON NOT NULL COMMENT '关键词',
    `catalogue` VARCHAR(255) DEFAULT NULL COMMENT '目录',
    `cover` VARCHAR(255) DEFAULT NULL COMMENT '封面',
    `inventory` INT NOT NULL DEFAULT 0 COMMENT '库存',
    `supplier_ids` JSON DEFAULT NULL COMMENT '供应商ID',
    `location` VARCHAR(255) DEFAULT NULL COMMENT '位置',
    `subbook_ids` JSON DEFAULT NULL COMMENT '子书籍ID',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_name` (`name`),
    KEY `idx_publisher` (`publisher`),
    KEY `idx_catalogue` (`catalogue`)
);

CREATE TRIGGER after_book_update AFTER UPDATE
ON book FOR EACH ROW
BEGIN
    DECLARE record_count INT;
    IF (NEW.inventory < 10) THEN
        SELECT COUNT(*) INTO record_count FROM out_of_stock_record WHERE book_id = NEW.id;
        if (record_count = 0) THEN
            INSERT INTO out_of_stock_record(book_id, number) VALUES (NEW.id, 20);
        ELSE
            UPDATE out_of_stock_record SET number = number + 20 WHERE book_id = NEW.id;
        END IF;
    END IF;
END;