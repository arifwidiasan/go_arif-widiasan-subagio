// 1. Insert
// a. Insert 5 operators pada table operators.
db.createCollection('operators');
db.operators.insertOne({_id:1, name:"xl", created_at:ISODate(), updated_at:ISODate()})
db.operators.insertOne({_id:2, name:"telkomsel", created_at:ISODate(), updated_at:ISODate()})
db.operators.insertOne({_id:3, name:"three", created_at:ISODate(), updated_at:ISODate()})
db.operators.insertOne({_id:4, name:"indosat", created_at:ISODate(), updated_at:ISODate()})
db.operators.insertOne({_id:5, name:"smartfren", created_at:ISODate(), updated_at:ISODate()})

// b. Insert 3 product type.
db.product_types.insertMany([
    {_id:1, name:"pulsa", created_at:ISODate(), updated_at:ISODate()},
    {_id:2, name:"paket data", created_at:ISODate(), updated_at:ISODate()},
    {_id:3, name:"perdana", created_at:ISODate(), updated_at:ISODate()}
    ])

// c. Insert 2 product dengan product type id = 1, dan operators id = 3.
db.products.insertMany([
    {_id:1, product_type_id:1, operator_id:3, code: "A1", name:"pulsa a", status:"1", created_at:ISODate(), updated_at:ISODate()},
    {_id:2, product_type_id:1, operator_id:3, code: "A2", name:"pulsa b", status:"1", created_at:ISODate(), updated_at:ISODate()}
    ])

// d. Insert 3 product dengan product type id = 2, dan operators id = 1.
db.products.insertMany([
    {_id:3, product_type_id:2, operator_id:1, code: "B1", name:"paket a", status:"1", created_at:ISODate(), updated_at:ISODate()},
    {_id:4, product_type_id:2, operator_id:1, code: "B2", name:"paket b", status:"1", created_at:ISODate(), updated_at:ISODate()},
    {_id:5, product_type_id:2, operator_id:1, code: "B3", name:"paket c", status:"1", created_at:ISODate(), updated_at:ISODate()},
    ])

// e. Insert 3 product dengan product type id = 3, dan operators id = 4.
db.products.insertMany([
    {_id:6, product_type_id:3, operator_id:4, code: "C1", name:"perdana a", status:"1", created_at:ISODate(), updated_at:ISODate()},
    {_id:7, product_type_id:3, operator_id:4, code: "C2", name:"perdana b", status:"1", created_at:ISODate(), updated_at:ISODate()},
    {_id:8, product_type_id:3, operator_id:4, code: "C3", name:"perdana c", status:"1", created_at:ISODate(), updated_at:ISODate()},
    ])

// f. Insert product description pada setiap product.
db.product_descriptions.insertMany([
    {_id:1, description:"ini pulsa A"},
    {_id:2, description:"ini pulsa B"},
    {_id:3, description:"ini paket A"},
    {_id:4, description:"ini paket B"},
    {_id:5, description:"ini paket C"},
    {_id:6, description:"ini perdana A"},
    {_id:7, description:"ini perdana B"},
    {_id:8, description:"ini perdana C"}
    ])

// g. Insert 3 payment methods.
db.payment_method.insertMany([
    {_id:1, name:"transfer bank", status:"1", created_at:ISODate(), updated_at:ISODate()},
    {_id:2, name:"kartu debit/kredit", status:"1", created_at:ISODate(), updated_at:ISODate()},
    {_id:3, name:"e-wallet", status:"1", created_at:ISODate(), updated_at:ISODate()}
    ])

// h. Insert 5 user pada tabel user.
db.users.insertMany([
    {_id:1, name:"andi", status:"1", dob:"1999-01-01", gender:'M', created_at:ISODate(), updated_at:ISODate()},
    {_id:2, name:"budi", status:"1", dob:"1999-01-01", gender:'M', created_at:ISODate(), updated_at:ISODate()},
    {_id:3, name:"wati", status:"1", dob:"1999-01-01", gender:'F', created_at:ISODate(), updated_at:ISODate()},
    {_id:4, name:"joko", status:"1", dob:"1999-01-01", gender:'M', created_at:ISODate(), updated_at:ISODate()},
    {_id:5, name:"indah", status:"1", dob:"1999-01-01", gender:'F', created_at:ISODate(), updated_at:ISODate()}
    ])

// i. Insert 3 transaksi di masing-masing user.
db.transactions.insertMany([
    {_id:1, user_id:1, payment_method_id:1, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:2, user_id:1, payment_method_id:2, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:3, user_id:1, payment_method_id:3, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    ])

db.transactions.insertMany([
    {_id:4, user_id:2, payment_method_id:1, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:5, user_id:2, payment_method_id:2, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:6, user_id:2, payment_method_id:3, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    ])

db.transactions.insertMany([
    {_id:7, user_id:3, payment_method_id:1, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:8, user_id:3, payment_method_id:2, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:9, user_id:3, payment_method_id:3, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    ])

db.transactions.insertMany([
    {_id:10, user_id:4, payment_method_id:1, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:11, user_id:4, payment_method_id:2, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:12, user_id:4, payment_method_id:3, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    ])

db.transactions.insertMany([
    {_id:13, user_id:5, payment_method_id:1, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:14, user_id:5, payment_method_id:2, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:15, user_id:5, payment_method_id:3, status:"sukses", total_qty:1, total_price:10000, created_at:ISODate(), updated_at:ISODate()},
    ])

// j. Insert 3 product di masing-masing transaksi.
db.transaction_detail.insertMany([
    {_id:1, product_id:1, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:2, product_id:2, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:3, product_id:3, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:4, product_id:4, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:5, product_id:5, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:6, product_id:6, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:7, product_id:7, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:8, product_id:8, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:9, product_id:1, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:10, product_id:2, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:11, product_id:3, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:12, product_id:4, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:13, product_id:5, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:14, product_id:6, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    {_id:15, product_id:7, status:"sukses", qty:1, price:10000, created_at:ISODate(), updated_at:ISODate()},
    ])

// 2. Read
// a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
db.users.find({gender:'M'})

// b. Tampilkan product dengan id = 3.
db.products.find({_id:3})

// c. Hitung jumlah user / pelanggan dengan status gender Perempuan.
db.users.find({gender:'F'}).count()

// d. Tampilkan data pelanggan dengan urutan sesuai nama abjad
db.users.find().sort({name: 1})

// e. Tampilkan 5 data pada data product
db.products.find().limit(5)

// 3. Update
// a. Ubah data product id 1 dengan nama ‘product dummy’.
db.products.updateMany({_id:1},
 { $set: {name:"product dummy", updated_at:ISODate()}})

// b. Update qty = 3 pada transaction detail dengan product id 1.
db.transaction_detail.updateMany({product_id:1},
 { $set: 
 {qty:3, updated_at:ISODate()}
 })

// 4 Delete
// a. Delete data pada tabel product dengan id 1.
db.products.deleteOne({_id:1})

// b. Delete pada pada tabel product dengan product type id 1.
db.products.deleteMany({product_type_id:1})
