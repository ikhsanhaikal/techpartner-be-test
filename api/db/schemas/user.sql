CREATE TABLE users (
  id   int  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255)    NOT NULL,
  password varchar(255)    NOT NULL,
	email varchar(255)   NOT NULL UNIQUE
);