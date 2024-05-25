CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE schemes
(
    id SERIAL PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);