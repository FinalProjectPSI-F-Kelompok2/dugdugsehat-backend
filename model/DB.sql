-- Database using PostgreSQL

CREATE TABLE users (
  email VARCHAR(100),
  pass VARCHAR(255) NOT NULL,
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

CREATE TABLE measure_type (
  type_id INT NOT NULL,
  type_name VARCHAR(100) NOT NULL,
  PRIMARY KEY(type_id)
);

CREATE TABLE health_data (
  email VARCHAR(100) NOT NULL,
  type_id INT NOT NULL,
  measure_date TIMESTAMP,
  measure_value INT,
  CONSTRAINT email FOREIGN KEY(email) REFERENCES users(email),
  CONSTRAINT type_id FOREIGN KEY(type_id) REFERENCES measure_type(type_id)
);

INSERT INTO measure_type VALUES (0, 'ecg'), (1, 'hr');