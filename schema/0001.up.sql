CREATE TABLE users
(
    uuid varchar(10) NOT NULL UNIQUE PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);