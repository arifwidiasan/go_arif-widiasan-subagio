# (4) Basic Programming

## Resume
Dalam materi ini, yang dipelajari adalah :
1. Golang
2. Variable
3. Branching & Looping

### Golang
Golang adalah salah satu bahasa backend yang sedang popular saat ini, golang merupakan Bahasa pemrograman open source yang dikembangkan oleh google. Bahasa ini dikembangkan dengan tujuan untuk mempermudah proses development namun tetap simple,handal, dan efisien. Golang merupakan Bahasa yang bagus untuk menulis low-level program yang bisa digunakan untuk membuat program aplikasi dan program system.

### Variable
Variable digunakan untuk menyimpan informasi pada program komputer dimana variable memberikan labelling data dengan deskripsi nama dan memiliki tipe data seperti Boolean,numeric, dan string. Golang juga ada operasi menggunakan operator aritmatika, Boolean, dan bitwise.

### Branching & Looping
Branching pada golang ada if else dan switch, untuk perulangan pada umunya menggunakan for, for juga bisa digunakan untuk perulangan while. Pada perulangan for juga bisa diberikan continue dan break dimana continue bertujuan untuk meneruskan perulangan dan break untuk menghentikan perulangan.

## Task
### 1. Menghitung Luas Permukaan Tabung
Buatlah sebuah program untuk menghitung luas permukaan tabung !. setelah program berhasil dibuat, coba lakukan improvement dengan merubah agar inputan menjadi scanf !.
> input : T = 20, r = 4  
output : 602.88

Berikut kode dari task ini :

[luasTabung.go](./praktikum/1_luasTabung/luasTabung.go)

Hasil kode program :

<img src="./screenshots/1_luas_tabung.jpg" width="900">

### 2. Grade Nilai
Seorang pengajar sedang memeriksa ujian siswa dan akan memberikan deskripsi nilai dari A-E dengan ketentuan sebagai berikut : 
> Nilai 80-100 : A  
Nilai 65-79 : B  
Nilai 50-64 : C  
Nilai 35-49 : D  
Nilai 0-34 : E

Tampilkan deskripsi Nilai dan Nama siswa saat pengajar tersebut memasukkan nilai dan nama yang dia inginkan.

Berikut kode dari task ini :

[gradeNilai.go](./praktikum/2_gradeNilai/gradeNilai.go)

Hasil kode program :

<img src="./screenshots/2_grade_nilai.jpg" width="900">

### 3. Faktor Bilangan
pada task ini, buatlah program untuk mencetak faktor sebuah bilangan.
> input : 20  
output :  
1  
2  
4  
5  
10  
20

Berikut kode dari task ini :

[faktorBilangan.go](./praktikum/3_faktorBilangan/faktorBilangan.go)

Hasil kode program :

<img src="./screenshots/3_faktor_bilangan.jpg" width="900">

### 4. Bilangan Prima
pada task ini, buatlah fungsi untuk menentukan bahwa sebuah bilangan termasuk bilangan prima atau tidak.
> fmt.Println(primeNumber(11)) // true  
fmt.Println(primeNumber(13)) // true  
fmt.Println(primeNumber(17)) // true  
fmt.Println(primeNumber(20)) // false  
fmt.Println(primeNumber(35)) // false  
fmt.Println(primeNumber(1)) // false

Berikut kode dari task ini :

[prima.go](./praktikum/4_prima/prima.go)

Hasil kode program :

<img src="./screenshots/4_prima.jpg" width="900">

### 5. Palindrome
kata palindrome adalah sebuah kata yang jika dibalik, tetap sama. contoh 'katak' dibalik tetaplah 'katak'. pada task ini, buatlah program untuk mendeteksi sebuah string termasuk palindrome atau tidak !
> fmt.Println(palindrome("civic")) // true  
fmt.Println(palindrome("katak")) // true  
fmt.Println(palindrome("kasur rusak")) // true  
fmt.Println(palindrome("mistar")) // false  
fmt.Println(palindrome("lion")) // false

Berikut kode dari task ini :

[palindrome.go](./praktikum/5_palindrome/palindrome.go)

Hasil kode program :

<img src="./screenshots/5_palindrome.jpg" width="900">

### 6. Exponentiation
pada task ini, buatlah fungsi untuk menghitung pangkat bilangan (contoh: x^n)
> fmt.Println(pangkat(2, 3)) // 8  
fmt.Println(pangkat(5, 3)) // 125  
fmt.Println(pangkat(10, 2)) // 100  
fmt.Println(pangkat(2, 5)) // 32  
fmt.Println(pangkat(7, 3)) // 343

Berikut kode dari task ini :

[pangkat.go](./praktikum/6_pangkat/pangkat.go)

Hasil kode program :

<img src="./screenshots/6_pangkat.jpg" width="900">