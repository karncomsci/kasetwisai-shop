CREATE TABLE users (
	id serial PRIMARY KEY,
	Fristname VARCHAR(255) NOT NULL,
	Lastname VARCHAR(255) NOT NULL,
	username VARCHAR ( 200 ) UNIQUE NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	password VARCHAR ( 200 ) NOT NULL,
	role VARCHAR(255) NOT NULL,
	provider VARCHAR(100) NOT NULL,
	photo    VARCHAR(255) NOT NULL,
	verified Boolean NOT NULL,
	status VARCHAR(2) NOT NULL,
	createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL
);