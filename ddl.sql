CREATE DATABASE `Account_Service_DB`;

USE Account_Service_DB;

CREATE TABLE users(
id VARCHAR(100) PRIMARY KEY NOT NULL,
username VARCHAR(100) NOT NULL,
email VARCHAR(100) NOT NULL,
password VARCHAR(100) NOT NULL,
phone_number INT NOT NULL unique,
date_of_birth DATE,
address VARCHAR(100),
balance INT NOT NULL,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
deleted_at DATETIME
);

CREATE TABLE top_up(
id VARCHAR(100) PRIMARY KEY NOT NULL,
user_id VARCHAR(100) NOT NULL,
amount INT NOT NULL,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
CONSTRAINT fk_top_up_user FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE transfer(
id VARCHAR(100) PRIMARY KEY NOT NULL,
user_id_sender VARCHAR(100) NOT NULL,
user_id_recipient VARCHAR(100) NOT NULL,
amount INT NOT NULL,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
CONSTRAINT fk_transfer_sender FOREIGN KEY (user_id_sender) REFERENCES users(id),
CONSTRAINT fk_transfer_recipient FOREIGN KEY (user_id_recipient) REFERENCES users(id)
);