DROP DATABASE IF EXISTS school_project;

CREATE DATABASE school_project;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
);

CREATE TABLE word_sets (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  words JSON NOT NULL,
  user_id INT NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);