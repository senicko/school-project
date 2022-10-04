DROP DATABASE IF EXISTS school_project;

CREATE DATABASE school_project;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  password BYTEA NOT NULL
);