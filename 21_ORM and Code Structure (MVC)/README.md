# (21) ORM and Code Structure (MVC)

## Resume
Dalam materi ini, yang dipelajari adalah :
1. GORM
2. Database Migration
3. Code Structure (MVC)

### GORM
Object Relational Mapping (ORM) adalah teknik programming untuk mengkonversi data yang tipe sistem tidak kompatibel menggunakan bahasa pemrograman berorientasi objek, GORM sendiri adalah ORM library untuk Golang. GORM merepresentasikan tabel database menjadi object struct. Keuntungan dari ORM adalah less repetitive query, otomatis mengambil data menjadi objek yang siap digunakan, cara mudah untuk screening data, dan beberapa ada cache query. Kelemahan dari ORM adalah menambah layer pada kode, memuat relation yang tidak perlu, complex query akan terlalu panjang untuk ditulis, spesifik sql function ada yang tidak support.

### Database Migration
Pada GORM ada fitur Database migration yang mana merupakan cara untuk update versi database untuk memenuhi perubahan versi aplikasi, perubahan bisa upgrade ke versi terbaru atau rollback ke versi sebelumnya. Keuntungan dari penggunaan database migration adalah memudahkan update dan rollback database, melacak perubahan pada struktur database, riwayat struktur database ditulis di kode, dan selalu compatible dengan perubahan versi aplikasi.

### Code Structure
Code Structure bertujuan untuk merestruktur project kita, struktur kode yang biasanya diterapkan adalah Model-View-Controller (MVC). MVC adalah cara yang populer untuk menata project kita, gambaran dari MVC adalah setiap bagian dari kode memiliki tujuan dan tujuan tersebut berbeda-beda. Kita membutuhkan struktur kode untuk mendapatkan modular application, implementasi pemisahan perhatian, dan mengurangi conflict saat melakukan versioning.

## Task
### 1. API CRUD User Using Database
pada task ini, Buat API CRUD User menggunakan database dengan spesifikasi seperti berikut!

| Route | HTTP | Description |
| --- | --- | --- |
| `/users` | GET | Get all users data |
| `/users/:id` | GET | Get single user by id |
| `/users` | POST | Create new user |
| `/users/:id` | DELETE | Delete user by id |
| `/users/id` | PUT | Update user information |

Kode Program dan hasil output nya adalah :

[main.go](./praktikum/problem_1/project/main.go)

1. Get all users data
   <br><br><img src="./screenshots/problem_1/getall.jpg" width="700">

2. Get single user by id
   <br><br><img src="./screenshots/problem_1/get1.jpg" width="700">

3. Create new user
   <br><br><img src="./screenshots/problem_1/create.jpg" width="700">

4. Delete user by id
   <br><br><img src="./screenshots/problem_1/delete.jpg" width="700">
   <br><br><img src="./screenshots/problem_1/deleteafter.jpg" width="700">

5. Update user information
   <br><br><img src="./screenshots/problem_1/update.jpg" width="700">
   <br><br><img src="./screenshots/problem_1/updateafter.jpg" width="700">

### 2.  Structuring Project with Layered Architecture
pada task ini, Buat API CRUD User menggunakan database dan terapkan code structure dengan spesifikasi seperti berikut!

| Route | HTTP | Description |
| --- | --- | --- |
| `/users` | GET | Get all users data |
| `/users/:id` | GET | Get single user by id |
| `/users` | POST | Create new user |
| `/users/:id` | DELETE | Delete user by id |
| `/users/id` | PUT | Update user information |

Kode Program dan hasil output nya adalah :

[main.go](./praktikum/problem_2/project/main.go)
[config.go](./praktikum/problem_2/project/config/config.go)
[user_controller.go](./praktikum/problem_2/project/controller/user_controller.go)
[user.go](./praktikum/problem_2/project/model/user.go)
[route.go](./praktikum/problem_2/project/route/route.go)

1. Get all users data
   <br><br><img src="./screenshots/problem_2/user_getall.jpg" width="700">

2. Get single user by id
   <br><br><img src="./screenshots/problem_2/user_get1.jpg" width="700">

3. Create new user
   <br><br><img src="./screenshots/problem_2/user_create.jpg" width="700">

4. Delete user by id
   <br><br><img src="./screenshots/problem_2/user_delete.jpg" width="700">
   <br><br><img src="./screenshots/problem_2/user_deleteafter.jpg" width="700">

5. Update user information
   <br><br><img src="./screenshots/problem_2/user_update.jpg" width="700">
   <br><br><img src="./screenshots/problem_2/user_updateafter.jpg" width="700">

| Route | HTTP | Description |
| --- | --- | --- |
| `/books` | GET | Get all books data |
| `/books/:id` | GET | Get single books by id |
| `/books` | POST | Create new books |
| `/books/:id` | DELETE | Delete books by id |
| `/books/id` | PUT | Update books information |

Kode Program dan hasil output nya adalah :

[main.go](./praktikum/problem_2/project/main.go)
[config.go](./praktikum/problem_2/project/config/config.go)
[book_controller.go](./praktikum/problem_2/project/controller/book_controller.go)
[book.go](./praktikum/problem_2/project/model/book.go)
[route.go](./praktikum/problem_2/project/route/route.go)

1. Get all books data
   <br><br><img src="./screenshots/problem_2/book_getall.jpg" width="700">

2. Get single book by id
   <br><br><img src="./screenshots/problem_2/book_get1.jpg" width="700">

3. Create new book
   <br><br><img src="./screenshots/problem_2/book_create.jpg" width="700">

4. Delete book by id
   <br><br><img src="./screenshots/problem_2/book_delete.jpg" width="700">
   <br><br><img src="./screenshots/problem_2/book_deleteafter.jpg" width="700">

5. Update book information
   <br><br><img src="./screenshots/problem_2/book_update.jpg" width="700">
   <br><br><img src="./screenshots/problem_2/book_updateafter.jpg" width="700">