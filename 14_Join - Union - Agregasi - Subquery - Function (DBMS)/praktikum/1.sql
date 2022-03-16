SET sql_mode = '';

ALTER TABLE `transaction_detail` ADD UNIQUE `unique_index`(`transaction_id`, `product_id`);

-- 1 --
-- a. Insert 5 operators pada table operators.
INSERT INTO operators (id,NAME) VALUES (1,'xl');
INSERT INTO operators (id,NAME) VALUES (2,'telkomsel');
INSERT INTO operators (id,NAME) VALUES (3,'3');
INSERT INTO operators (id,NAME) VALUES (4,'indosat');
INSERT INTO operators (id,NAME) VALUES (5,'smartfren');

SELECT * FROM operators;

-- b. Insert 3 product type.
INSERT INTO product_types (id,NAME) VALUES (1,'pulsa');
INSERT INTO product_types (id,NAME) VALUES (2,'paket data');
INSERT INTO product_types (id,NAME) VALUES (3,'prabayar');

SELECT * FROM product_types;

-- c. Insert 2 product dengan product type id = 1, dan operators id = 3.
INSERT INTO products (id,product_type_id,operator_id,CODE,NAME,STATUS) VALUES (1,1,3,'A1','Pulsa A','100');
INSERT INTO products (id,product_type_id,operator_id,CODE,NAME,STATUS) VALUES (2,1,3,'A1','Pulsa B','100');

SELECT * FROM products;

-- d. Insert 3 product dengan product type id = 2, dan operators id = 1.
INSERT INTO products (id,product_type_id,operator_id,CODE,NAME,STATUS) VALUES (3,2,1,'B1','paket A','100');
INSERT INTO products (id,product_type_id,operator_id,CODE,NAME,STATUS) VALUES (4,2,1,'B2','paket B','100');
INSERT INTO products (id,product_type_id,operator_id,CODE,NAME,STATUS) VALUES (5,2,1,'B3','paket C','100');

SELECT * FROM products;

-- e. Insert 3 product dengan product type id = 3, dan operators id = 4.
INSERT INTO products (id,product_type_id,operator_id,CODE,NAME,STATUS) VALUES (6,3,4,'C1','perdana A','100');
INSERT INTO products (id,product_type_id,operator_id,CODE,NAME,STATUS) VALUES (7,3,4,'C2','perdana B','100');
INSERT INTO products (id,product_type_id,operator_id,CODE,NAME,STATUS) VALUES (8,3,4,'C3','perdana C','100');

SELECT * FROM products;

-- f. Insert product description pada setiap product.
INSERT INTO product_descriptions (id,DESCRIPTION) VALUES (1,'ini pulsa A');
INSERT INTO product_descriptions (id,DESCRIPTION) VALUES (2,'ini pulsa B');
INSERT INTO product_descriptions (id,DESCRIPTION) VALUES (3,'ini paket A');
INSERT INTO product_descriptions (id,DESCRIPTION) VALUES (4,'ini paket B');
INSERT INTO product_descriptions (id,DESCRIPTION) VALUES (5,'ini paket C');
INSERT INTO product_descriptions (id,DESCRIPTION) VALUES (6,'ini perdana A');
INSERT INTO product_descriptions (id,DESCRIPTION) VALUES (7,'ini perdana B');
INSERT INTO product_descriptions (id,DESCRIPTION) VALUES (8,'ini perdana C');

SELECT * FROM product_descriptions;

-- g. Insert 3 payment methods.
INSERT INTO payment_method (id,NAME,STATUS) VALUES (1,'transfer bank','100');
INSERT INTO payment_method (id,NAME,STATUS) VALUES (2,'kartu debit/kredit','100');
INSERT INTO payment_method (id,NAME,STATUS) VALUES (3,'e-wallet','100');

SELECT * FROM payment_method;

-- h. Insert 5 user pada tabel user.
INSERT INTO users (id,NAME,STATUS,dob,gender) VALUES (1,'andi','1','1999-01-01','M');
INSERT INTO users (id,NAME,STATUS,dob,gender) VALUES (2,'budi','1','1998-11-11','M');
INSERT INTO users (id,NAME,STATUS,dob,gender) VALUES (3,'wati','1','1997-04-14','F');
INSERT INTO users (id,NAME,STATUS,dob,gender) VALUES (4,'joko','1','1998-09-23','M');
INSERT INTO users (id,NAME,STATUS,dob,gender) VALUES (5,'indah','1','1998-12-30','F');

SELECT * FROM users;

-- i. Insert 3 transaksi di masing-masing user.
INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (1,1,1,'sukses',1,10000);
INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (2,1,2,'sukses',1,11000);
INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (3,1,3,'sukses',1,12000);

INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (4,2,1,'sukses',1,10000);
INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (5,2,2,'sukses',1,11000);
INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (6,2,3,'sukses',1,12000);

INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (7,3,1,'sukses',1,10000);
INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (8,3,2,'sukses',1,11000);
INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (9,3,3,'sukses',1,12000);

INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (10,4,1,'sukses',1,10000);
INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (11,4,2,'sukses',1,11000);
INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (12,4,3,'sukses',1,12000);

INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (13,5,1,'sukses',1,10000);
INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (14,5,2,'sukses',1,11000);
INSERT INTO transactions (id,user_id,payment_method_id,STATUS,total_qty,total_price) VALUES (15,5,3,'sukses',1,12000);

SELECT * FROM transactions;

-- j. Insert 3 product di masing-masing transaksi.
INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (1,1,'sukses',1,10000);
INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (2,2,'sukses',1,11000);
INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (3,3,'sukses',1,12000);

INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (4,4,'sukses',1,10000);
INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (5,5,'sukses',1,11000);
INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (6,6,'sukses',1,12000);

INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (7,7,'sukses',1,10000);
INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (8,8,'sukses',1,11000);
INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (9,1,'sukses',1,12000);

INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (10,1,'sukses',1,10000);
INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (11,2,'sukses',1,11000);
INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (12,3,'sukses',1,12000);

INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (13,1,'sukses',1,10000);
INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (14,2,'sukses',1,11000);
INSERT INTO transaction_detail (transaction_id,product_id,STATUS,qty,price) VALUES (15,3,'sukses',1,12000);

SELECT * FROM transaction_detail;

-- 2 --
-- a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT * FROM users WHERE gender='M';

-- b. Tampilkan product dengan id = 3.
SELECT * FROM products WHERE id=3;

-- c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * FROM users WHERE created_at > CURRENT_DATE - INTERVAL 7 DAY AND NAME LIKE '%a%';

-- d. Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(gender) FROM users WHERE gender='F';

-- e. Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM users ORDER BY NAME ASC;

-- f. Tampilkan 5 data pada data product
SELECT * FROM products LIMIT 5;

-- 3 --
-- a. Ubah data product id 1 dengan nama ‘product dummy’.
UPDATE products SET NAME = 'product dummy' WHERE id = 1;
SELECT * FROM products;

-- b. Update qty = 3 pada transaction detail dengan product id 1.
UPDATE transaction_detail SET qty = 3 WHERE product_id = 1;
SELECT * FROM transaction_detail;

-- 4 --
-- a. Delete data pada tabel product dengan id 1.
DELETE FROM products WHERE id=1;
SELECT * FROM products;

-- b. Delete pada pada tabel product dengan product type id 1.
DELETE FROM products WHERE product_type_id=1;
SELECT * FROM products;
