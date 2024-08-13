CREATE TABLE categories (
  id   int  NOT NULL AUTO_INCREMENT PRIMARY KEY,
	nama varchar(255) NOT NULL,
	tipe enum("pemasukan", "pengeluaran") NOT NULL,
	deskripsi text
);