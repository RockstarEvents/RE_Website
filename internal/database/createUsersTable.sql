CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    status INT NOT NULL
);


CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    shape VARCHAR(255) NOT NULL,
    place VARCHAR(255) NOT NULL,
    begin_time TIMESTAMP NOT NULL,
    duration VARCHAR(255) NOT NULL
);


CREATE VIEW contacts AS
SELECT id, username, email
FROM users;

