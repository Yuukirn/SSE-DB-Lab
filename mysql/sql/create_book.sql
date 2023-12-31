USE bookstore_management_system;

DROP TABLE IF EXISTS book;

CREATE TABLE book (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL DEFAULT '',
    authors JSON NOT NULL,
    publisher VARCHAR(255) NOT NULL DEFAULT '',
    price DOUBLE NOT NULL DEFAULT 0.0,
    keywords JSON NOT NULL,
    catalogue VARCHAR(255) DEFAULT NULL,
    cover VARCHAR(255) DEFAULT NULL,
    inventory INT NOT NULL DEFAULT 0,
    supplier_ids JSON DEFAULT NULL,
    location VARCHAR(255) DEFAULT NULL,
    subbook_ids JSON DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_name ON book (name);
CREATE INDEX idx_publisher ON book (publisher);
CREATE INDEX idx_catalogue ON book (catalogue);

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