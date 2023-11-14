-- Database using PostgreSQL

CREATE TABLE users (
  email VARCHAR(100),
  password VARCHAR(255) NOT NULL,
  PRIMARY KEY(email)
);

CREATE TABLE profile (
  email VARCHAR(100) NOT NULL,
  full_name VARCHAR(255) NOT NULL,
  body_height INT,
  body_weight INT,
  age INT,
  sex BOOLEAN,
  CONSTRAINT email FOREIGN KEY(email) REFERENCES users(email)
);

CREATE TABLE health_data (
  email VARCHAR(100) NOT NULL,
  measure_date TIMESTAMP,
  ecg FLOAT(2) NOT NULL,
  heart_rate INT NOT NULL,
  CONSTRAINT email FOREIGN KEY(email) REFERENCES users(email)
);