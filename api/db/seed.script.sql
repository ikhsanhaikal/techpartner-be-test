INSERT INTO users (name, email, password) 
VALUES 
  ('Budi', 'budi@mail.com', 'passwordnyabudi'), 
  ('Agus', 'agus@mail.com', 'passwordnyaagus'), 
  ('Ani', 'ani@mail.com', 'passwordnyaani');

INSERT INTO accounts (saldo, user_id) 
VALUES 
  (5000000, 1), 
  (3500000, 2), 
  (10000000, 3);


INSERT INTO categories (nama, tipe) 
VALUES 
  ('gaji', 'pemasukan'), 
  ('tunjangan', 'pemasukan'), 
  ('bonus', 'pemasukan'),
  ('sewa kost', 'pengeluaran'),
  ('makan', 'pengeluaran'),
  ('pakaian', 'pengeluaran');

UPDATE accounts
SET saldo = saldo - 350000
WHERE id = 1;

UPDATE accounts
SET saldo = saldo - 1500000
WHERE id = 3;

INSERT INTO transactions (user_id, rek_id, kategori_id,
nominal, deskripsi) 
VALUES 
  (1, 1, 6, 350000, "kemeja uniqlo promo"), 
  (3, 3, 4, 1500000, "byr kost bulan kemarin");

