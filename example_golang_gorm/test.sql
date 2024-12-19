INSERT INTO users (first_name, last_name, email, password) VALUES
('Do', 'Gia Huy', 'dogiahuy@example.com', MD5('password001')),
('Tran', 'Anh Thu', 'trânnhthu@example.com', MD5('password002')),
('Do Nguyen', 'Ha Nam', 'donguyenhanam@example.com', MD5('password003'));

INSERT INTO projects (name, project_started_at, user_id) VALUES
('Project A', NOW(), 1),
('Project B', NOW(), 2),
('Project C', NOW(), 3),
('Project D', NOW(), 1),
('Project E', NOW(), 2),
('Project F', NOW(), 3),
('Project G', NOW(), 1),
('Project H', NOW(), 2),
('Project I', NOW(), 3),
('Project J', NOW(), 1);

SELECT * FROM users;

SELECT * FROM projects;
