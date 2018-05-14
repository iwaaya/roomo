CREATE DATABASE roomo;

USE roomo;

CREATE TABLE collections ( 
    id INTEGER NOT NULL AUTO_INCREMENT,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    location VARCHAR(255) NOT NULL,
    primary key(id)
);
