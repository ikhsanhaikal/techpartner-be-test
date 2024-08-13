SET FOREIGN_KEY_CHECKS = 0;
drop table if exists users;
drop table if exists categories;
drop table if exists transactions;
drop table if exists accounts;
SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE categories (
  id   int  NOT NULL AUTO_INCREMENT PRIMARY KEY,
	nama varchar(255),
	tipe enum("pemasukan", "pengeluaran"),
	deskripsi text
);

CREATE TABLE users (
  id   int  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255)    NOT NULL,
  password varchar(255)    NOT NULL,
	email varchar(255)   NOT NULL UNIQUE
);

CREATE TABLE accounts (
  id   int  NOT NULL AUTO_INCREMENT PRIMARY KEY,
	saldo decimal(10, 2) NOT NULL,
	user_id int NOT NULL,
	created_at  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
	FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE transactions (
  id   int  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  kategori_id int NOT NULL,
	rek_id int NOT NULL,
	user_id int NOT NULL,
	nominal decimal NOT NULL,
	deskripsi text,
	created_at  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
	deleted_at timestamp,
	FOREIGN KEY (kategori_id) REFERENCES categories(id),
	FOREIGN KEY (rek_id) REFERENCES accounts(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);