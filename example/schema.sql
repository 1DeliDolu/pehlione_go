-- Create database
CREATE DATABASE IF NOT EXISTS my_go_database;
USE my_go_database;

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO users (password, username, email) VALUES
    ('password123', 'john_doe', 'john@pehlione.com'),
    ('password456', 'jane_doe', 'jane@pehlione.com'),
    ('password789', 'bob_smith', 'bob@pehlione.com');