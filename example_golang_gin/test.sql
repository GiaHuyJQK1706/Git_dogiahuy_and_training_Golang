INSERT INTO users (first_name, last_name, email, password, created_at, updated_at)
VALUES 
('John', 'Doe', 'john.doe@example.com', 'password123', NOW(), NOW()),
('Jane', 'Smith', 'jane.smith@example.com', 'password456', NOW(), NOW()),
('Alice', 'Johnson', 'alice.johnson@example.com', 'password789', NOW(), NOW()),
('Do', 'Gia Huy', 'huydogiaccac@gmail.com', 'abcd1234', NOW(), NOW());

SELECT * FROM users; -- Kiem tra bang