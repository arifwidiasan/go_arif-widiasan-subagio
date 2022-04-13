# (26) Clean Architecture Unit Test

## Resume
Dalam materi ini, yang dipelajari adalah :
1. Clean Architecture
2. CA Layer
3. Unit Test

### Clean Architecture
Constraint sebelum membangun clean architecture adalah Independent of framework, Testable, Independent of UI, Independent of Database, Independent of any External. Keuntungan dari clean architecture adalah :
- struktur yang standar sehingga mudah diterapkan pada project.
- faster development dalam jangka waktu yang panjang.
- mocking dependencies menjadi hal yang trivial pada unit test.
- mudah berganti dari prototype ke solusi yang sebenarnya.

### CA Layer
Kita dapat mendefinisikan dasar dari 3-layer architecture (bisa lebih layer) yaitu :
- Entity Layer : Object yang merefleksikan konsep dari managemen aplikasi (opsional)
- Use Case - Domain layer : mengandung business logic.
- Controller - Presentation Layer : menampilkan atau representasi data ke layar dan menghandle user interaction.
- Drivers - Data Layer : Memanage data aplikasi.

Domain Driver Design (DDD) adalah sebuah pendekatan untuk mengembangkan software kompleks yang menghubungkan konsep bisnis inti dan implementasi teknikal secara mendalam. Clean architecture merupakan arsitektur software sedangkan DDD merupakan teknik software design

### Unit Test
Software testing adalah proses analisa dari item software untuk memeriksa perbedaan dari kondisi sekarang dengan kondisi yang diperlukan dan untuk mengevaluasi fitur yang ada pada software. Karena sudah diterapkan clean architecture, maka kdapat menerapkan testing pada setiap layer seeprti use case, controller, dan driver. Untuk mengetahui result dan expected result dapat menggunakan assert dan untuk membuat database sementara bisa melakukan mocking.

## Task
### 1. Rewrite
Pada task ini, Setelah berhasil mengubah project di bawah ini menjadi Clean Architecture, tambahkan unit test untuk tiap fungsi yang ada!

[https://github.com/hadihammurabi/belajar-go-echo](https://github.com/hadihammurabi/belajar-go-echo)

Unit test yang dibuat harus memenuhi:
- Code coverage lebih dari 80%
- Menerapkan mocking

Hasil dari task tersebut adalah :

**Masih kesulitan untuk menerapkan mocking dan unit test di clean architecture**