-- Dùng CSDL:
USE testdb;

-- Tạo các bảng:
CREATE TABLE `categories` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `description` text NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `items` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `amount` int(10) NOT NULL,
    `category_id` int(10) NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tạo dữ liệu giả để test:
INSERT INTO categories (name, description) VALUES
    ('Electronics', 'Devices and gadgets'),
    ('Books', 'Printed and digital books'),
    ('Clothing', 'Men and women apparel'),
    ('Groceries', 'Daily food and household items'),
    ('Toys', 'Toys and games for children');

INSERT INTO items (name, amount, category_id) VALUES
    ('Smartphone', 50, 1),
    ('Laptop', 20, 1),
    ('Headphones', 100, 1),
    ('Fiction Book', 75, 2),
    ('Science Book', 30, 2),
    ('Jeans', 40, 3),
    ('T-Shirt', 60, 3),
    ('Bread', 200, 4),
    ('Milk', 150, 4),
    ('Action Figure', 25, 5),
    ('Board Game', 10, 5),
    ('Puzzle', 15, 5);

-- Câu 1: Lấy thông tin categories và số lượng items của mỗi category
SELECT 
    c.id AS category_id,
    c.name AS category_name,
    COUNT(i.id) AS item_count
FROM 
    categories c
LEFT JOIN 
    items i 
ON 
    c.id = i.category_id
GROUP BY 
    c.id, c.name;

-- Câu 2: Lấy thông tin categories và tổng amount của items trong category
SELECT 
    c.id AS category_id,
    c.name AS category_name,
    COALESCE(SUM(i.amount), 0) AS total_amount
FROM 
    categories c
LEFT JOIN 
    items i 
ON 
    c.id = i.category_id
GROUP BY 
    c.id, c.name;

-- Câu 3: Lấy thông tin categories có ít nhất 1 item có amount > 40
SELECT 
    c.id AS category_id,
    c.name AS category_name
FROM 
    categories c
INNER JOIN 
    items i 
ON 
    c.id = i.category_id
WHERE 
    i.amount > 40
GROUP BY 
    c.id, c.name;

-- Câu 4: Xóa những categories không có items
DELETE FROM categories
WHERE id NOT IN (SELECT DISTINCT category_id FROM items WHERE category_id IS NOT NULL);