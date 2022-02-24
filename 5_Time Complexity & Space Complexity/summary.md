# (5) Time Complexity & Space Complexity

## Resume
Dalam materi ini, yang dipelajari adalah :
1. Time Complexity
2. Perbedaan Tiap Time Complexity
3. Space Complexity

### Time Complexity
Penggunaan time complexity memudahkan untuk memperkirakan waktu yang dibutuhkan untuk program tersebut berjalan. complexity dapat dilihat sebagai maximum number dari primitive opperation yang program tersebut mungkin dieksekusi, operasi atau operator tersebut dinamakan dominan. complexity dilambangkan dengan Big-O Notation.

### Perbedaan Tiap Time Complexity
ada banyak sekali jenis time complexity, antara lain adalah :
- Constant Time : O(i)
- Linear Time : O(n) atau O(n+m)
- Logarithmic Time : O(log n)
- Quadratic Time : O(n^2)
setiap time complexity memiliki waktu eksekusi yang berbeda-beda dan dapat dilihat pada gambar berikut

<img src="https://i.imgur.com/EoBwqWm.jpeg">

### Space Complexity
Memory limit memberikan informasi untuk memperkirakan space complexity. kita bisa memperkirakan banyak variabel yang kamu deklarasikan pada program, dengan kata lain jika kita memiliki angka konstan variabel maka kita juga mempunyai space complexity yang konstan

## Task
### 1. Bilangan Prima
pada task ini, buatlah program untuk menentukan bilangan tersebut termasuk bilangan prima atau tidak dengan kompleksitas lebih cepat dari O(n) atau O(n/2)
>input : 1000000007  
Output: Bilangan Prima

Berikut kode dari task ini :

[prime.go](./praktikum/1_primeNumber/prime.go)

Hasil kode program :

<img src="./screenshots/1_prima.jpg" width="900">

### 1. Fast Exponentiation
pada task ini, buatlah fungsi program untuk menghitung x^n dengan kompleksitas lebih optimal dari O(n) yaitu logarithmic time (O(log n))
>input : x = 2, n = 3  
Output: 8
>fmt.Println(pow(2, 3))  // 8
fmt.Println(pow(5, 3))  // 125
fmt.Println(pow(10, 2)) // 100
fmt.Println(pow(2, 5))  // 32
fmt.Println(pow(7, 3))  // 343


Berikut kode dari task ini :

[fastExp.go](./praktikum/2_fastExponentiation/fastExp.go)

Hasil kode program :

<img src="./screenshots/2_fast_exp.jpg" width="900">