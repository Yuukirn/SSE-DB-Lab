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
    publisher_ids JSON DEFAULT NULL,
    location VARCHAR(255) DEFAULT NULL,
    subbook_ids JSON DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_name ON book (name);
CREATE INDEX idx_publisher ON book (publisher);
