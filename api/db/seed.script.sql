INSERT INTO users (name, email, password) 
VALUES 
  ('Budi', 'budi@mail.com', 'passwordnyabudi'), 
  ('Agus', 'agus@mail.com', 'passwordnyaagus'), 
  ('Ani', 'ani@mail.com', 'passwordnyaani');


INSERT INTO categories (nama, tipe) 
VALUES 
  ('gaji', 'pemasukan'), 
  ('tunjangan', 'pemasukan'), 
  ('bonus', 'pemasukan'),
  ('sewa kost', 'pengeluaran'),
  ('makan', 'pengeluaran'),
  ('pakaian', 'pengeluaran');


INSERT INTO transactions (user_id, kategori_id,
nominal, deskripsi) 
VALUES 
  (1, 1, 5000000, ""), 
  (1, 6, 3500000, "kemeja uniqlo promo"), 
  (2, 1, 6000000, ""),
  (3, 1, 5500000, ""),
  (3, 4, 1500000, "byr kost bulan kemarin");

