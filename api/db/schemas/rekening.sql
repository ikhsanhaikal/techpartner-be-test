CREATE TABLE accounts (
  id   int  NOT NULL AUTO_INCREMENT PRIMARY KEY,
	saldo decimal(10, 2) NOT NULL,
	user_id int NOT NULL,
	created_at  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
	FOREIGN KEY (user_id) REFERENCES users(id)
);