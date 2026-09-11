CREATE DATABASE beverages_db;
USE beverages_db;

-- DDL
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role ENUM('admin', 'kepala_toko') NOT NULL
);

CREATE TABLE customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(30)
);

CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(12,2) NOT NULL,
    stock INT NOT NULL DEFAULT 100
);

CREATE TABLE payment_methods (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT NOT NULL,
    payment_method_id INT NOT NULL,
    subtotal DECIMAL(12,2) NOT NULL,
    tax DECIMAL(12,2) NOT NULL DEFAULT 0,
    discount DECIMAL(12,2) NOT NULL DEFAULT 0,
    total DECIMAL(12,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE transaction_details (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(12,2) NOT NULL,
    subtotal DECIMAL(12,2) NOT NULL,
    FOREIGN KEY (transaction_id)
        REFERENCES transactions(id)
        ON DELETE CASCADE,
    FOREIGN KEY (product_id)
        REFERENCES products(id)
);

-- DML
INSERT INTO payment_methods 
    (id, name)
VALUES
    (1, 'Cash'),
    (2, 'Debit'),
    (3, 'Credit Card'),
    (4, 'Transfer');


INSERT INTO products 
    (name, price, stock)
VALUES
    ('Espresso', 18000, 100),
    ('Americano', 20000, 100),
    ('Cappuccino', 25000, 100),
    ('Cafe Latte', 25000, 100),
    ('Mocha', 28000, 100),
    ('Caramel Macchiato', 30000, 100),
    ('Matcha Latte', 28000, 100),
    ('Chocolate', 24000, 100),
    ('Thai Tea', 22000, 100),
    ('Lemon Tea', 18000, 100),
    ('Milk Tea', 22000, 100),
    ('Strawberry Smoothie', 30000, 100),
    ('Mango Smoothie', 30000, 100),
    ('Mineral Water', 10000, 100),
    ('Iced Tea', 12000, 100);
