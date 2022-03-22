# (18) System Design

## Resume
Dalam materi ini, yang dipelajari adalah :
1. Diagram Design and Distributed System
2. Job Queue and Microservices
3. Cache and Indexing

### Diagram Design and Distributed System
Diagram adalah sebuah simbol representasi dari informasi menggunakan teknik visualisasi. Diagram design yang sering dibuat adalah flowchart, use case, entity relationship diagram, dan high level architecture. Setiap kali kita mendesain system yang besar, kita perlu mempertimbangkan berbagai hal seperti apa saja arsitektur berbeda yang bisa digunakan, bagaimana mereka bekerja satu sama lain, bagaimana pemanfaatan penggunaan arsitektur tersebut, sehingga kita familiar dengan pemikiran tersebut untuk mengerti konsep distributed system. Karakteristik kunci dari distributed system adalah
- Scability : kapabilitas sistem, process, atau network untuk berkembang.
- Reliability : kemungkinan ada error pada sistem.
- Avaibility : berapa lama sistem dapat beroperasi.
- Efficiency : tingkat efektifitas sistem.
- Serviceibility or Managebility : mudah tidaknya perbaikan atau manage sistem tersebut.

### Job Queue and Microservices
Pada system software, job queue adalah struktur data yang diurus oleh job scheduler software yang mengandung job untuk dikerjakan. Work queue adalah framework untuk membangun master-worker application yang menjangkau banyak mesin dari cluster, clouds, and grids. Monolithic application memiliki satu basis kode dengan banyak modul, microservices adalah layanan yang dapat digunakan secara independen yang bermodel sekitar business domain.

### Cache and Indexing
Cache digunakan pada requested data terbaru yang kemungkinan besar akan dipanggil lagi, cache digunakan hampir pada seluruh layer computing : hardware, operating system, web browsers, web application dan lain - lain. indexing adalah cara untuk optimasi performa database dengan minimize banyaknya penyimpanan yang diakses pada saat query diproses.

## Task
### 1. Diagram
pada task ini, ada kasus dimana akan dikembangkan sistem yang dapat digunakan untuk mencatat pengeluaran seseorang dalam jangka waktu harian

[File ERD](./praktikum/erd.drawio)

[File Use Case](./praktikum/usecase.drawio)

- Buat desain ERD
<br><br><img src="./screenshots/1a.jpg" width="700">

- Buat Use Case
<br><br><img src="./screenshots/1b.jpg" width="700">

### 2. Query
Terdapat sebuah Query dalam format SQL
```
SELECT * FROM users;
```
Dengan tujuan yang sama, tuliskan dalam bentuk perintah:
1. Redis
```
users * 
atau 
FT.SEARCH users
```

2. Neo4J
```
MATCH (n:users)
```

3. Cassandra
```
SELECT * FROM users;
```