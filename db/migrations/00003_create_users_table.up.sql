CREATE TABLE IF NOT EXISTS users (
	id serial PRIMARY KEY,
	"name" varchar(255) NULL,
	username varchar(255) NULL,
	email varchar(255) NOT NULL,
	password varchar(255) NOT NULL
);