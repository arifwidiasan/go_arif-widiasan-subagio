-- 1. Gabungkan data transaksi dari user id 1 dan user id 2.
SELECT * FROM transactions WHERE user_id=1
UNION
SELECT * FROM transactions WHERE user_id=2;

-- 2. Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(total_price) FROM transactions WHERE user_id=1;

-- 3. Tampilkan total transaksi dengan product type 2.
SELECT COUNT(1) FROM transaction_detail t INNER JOIN products p
ON t.product_id = p.id
WHERE p.product_type_id=2;

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT * FROM products p INNER JOIN product_types pt 
ON p.product_type_id = pt.id;

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user.
/* SELECT * FROM transactions t 
INNER JOIN transaction_detail td ON t.id = td.transaction_id
inner join products p on p.id=td.product_id
inner join users u on t.user_id=u.id; */

SELECT * FROM transactions
INNER JOIN products
INNER JOIN users;

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
DELIMITER $$
CREATE TRIGGER delete_transaction_detail
BEFORE DELETE ON transactions FOR EACH ROW
BEGIN
DECLARE v_trans_id INT;
SET v_trans_id = OLD.id;
DELETE FROM transaction_detail WHERE  transaction_id = v_trans_id;
END$$

DELETE FROM transactions WHERE id=3;

SELECT * FROM transaction_detail;

-- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
DELIMITER $$
CREATE TRIGGER decrease_qty
BEFORE DELETE ON transaction_detail FOR EACH ROW
BEGIN
DECLARE v_id INT;
DECLARE v_qty INT;
SET v_id = OLD.transaction_id;
SET v_qty = OLD.qty;
UPDATE transactions SET total_qty = total_qty - v_qty WHERE id = v_id;
END$$

DELETE FROM transaction_detail WHERE transaction_id=4;

SELECT * FROM transactions;

-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
SELECT * FROM products WHERE id NOT IN
(SELECT product_id FROM transaction_detail);