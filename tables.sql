-- Users table DDL
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Insert some random data
INSERT INTO users (name) VALUES ('Alice');
INSERT INTO users (name) VALUES ('Bob');
INSERT INTO users (name) VALUES ('Charlie');
INSERT INTO users (name) VALUES ('Diana');
INSERT INTO users (name) VALUES ('Evan');
INSERT INTO users (name) VALUES ('Fiona');
INSERT INTO users (name) VALUES ('George');
INSERT INTO users (name) VALUES ('Hannah');
INSERT INTO users (name) VALUES ('Ian');
INSERT INTO users (name) VALUES ('Julia');
