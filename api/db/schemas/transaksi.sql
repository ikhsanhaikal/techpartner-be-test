CREATE TABLE transactions (
  id   int  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  kategori_id int NOT NULL,
	user_id int NOT NULL,
	nominal decimal NOT NULL,
	deskripsi text,
	created_at  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
	deleted_at timestamp,
	FOREIGN KEY (kategori_id) REFERENCES categories(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);