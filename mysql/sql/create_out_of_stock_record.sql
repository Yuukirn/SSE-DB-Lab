USE bookstore_management_system;

DROP TABLE IF EXISTS `out_of_stack_record`;

CREATE TABLE `out_of_stock_record` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `book_id` INT NOT NULL DEFAULT 0,
    `number` INT NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`book_id`) REFERENCES `book`(`id`)
);

CREATE INDEX `book_id` ON `out_of_stock_record` (`book_id`);
