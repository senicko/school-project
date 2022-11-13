DROP DATABASE IF EXISTS school_project;

CREATE DATABASE school_project;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  jokes JSONB
);
