INSERT INTO tasks 
(id, title, description, priority, status, created_at, updated_at, updated_by, due_date, jira_url, category_id) 
VALUES 
('e92f3b7d-0a1c-42ee-8862-60fa6ddcc567', 'Task 1', 'Description for Task 1', 2, 'assigned', '2022-01-15', '2022-01-20', '9c8b7a6d-5e4f-4032-8cab-eaacbcde9f12', '2022-02-01', 'https://jira.example.com/task1', '97b54c0a-d6e8-4f3a-af10-3ab2cdd7eeb2');

INSERT INTO tasks 
(id, title, description, priority, status, created_at, updated_at, updated_by, due_date, jira_url, category_id) 
VALUES 
('5c9f0d4b-6a2e-487b-bd31-0ddab04ce738', 'Task 2', 'Description for Task 2', 1, 'in progress', '2022-02-01', '2022-02-05', '9c8b7a6d-5e4f-4032-8cab-eaacbcde9f12', '2022-02-15', 'https://jira.example.com/task2', '97b54c0a-d6e8-4f3a-af10-3ab2cdd7eeb2'); 

INSERT INTO users (id, username, password, email, created_at) VALUES ('9c8b7a6d-5e4f-4032-8cab-eaacbcde9f12', 'john_doe', 'password123', 'johndoe@example.com', '2023-10-05 10:30:15'); 

INSERT INTO categories (id, name, description, created_at, updated_at) VALUES ('97b54c0a-d6e8-4f3a-af10-3ab2cdd7eeb2', 'Electronics', 'A wide variety of electronic devices and accessories.', '2022-08-15 09:30:00', '2022-09-02 14:45:00');