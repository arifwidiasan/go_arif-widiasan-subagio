# (16) MongoDB : Advanced Query - Array - Aggregation

## Resume
Dalam materi ini, yang dipelajari adalah :
1. Advanced Query
2. Array
3. Aggregation

### Advanced Query
Pada mongoDB terdapat advanced query

### Array
yeet

### Aggregation
teey

## Task
### 1. Aggregation MongoDB
pada task ini kita membuat query aggregation pada MongoDB dengan perintah seperti berikut :

[File Daftar Query MongoDB](./praktikum/query.js)

1. Gabungkan (menampilkan) data buku dari author id 1 dan author id 2.
```javascript
db.books.find({ $or: [{authorID:1},{authorID:2}]})
```
<img src="./screenshots/1.jpg" width="600">

2. Tampilkan daftar buku dan harga author id 1.
```javascript
db.books.aggregate([
    {
        $match: {
            authorID: 1
        }
    },
    {
        $project: {
            _id: 1,
            title: 1,
            price: 1
        }
    }
])
```
<img src="./screenshots/2.jpg" width="600">

3. Tampilan total jumlah halaman buku author id 2.
```javascript
db.books.aggregate([
    {
        $match: {
            authorID: 2
        }
    },
    {
        $group: {
            totalPages: { $sum : "$stats.page"},
            _id: 2
        }
    }
])
```
<img src="./screenshots/3.jpg" width="600">

4. Tampilkan semua field books and authors terkait.
```javascript
db.authors.aggregate([
    {
        $lookup : {
            from: "books",
            localField: "_id",
            foreignField : "authorID",
            as : "books"
        }
    }
])
```
<img src="./screenshots/4a.jpg" width="600">

```javascript
db.books.aggregate([
    {
        $lookup : {
            from: "authors",
            localField: "authorID",
            foreignField : "_id",
            as : "authors"
        }
    }
])
```
<img src="./screenshots/4b.jpg" width="600">

5. Tampilkan semua field books, authors, dan publishers terkait.
```javascript
db.books.aggregate([
    {
        $lookup : {
            from: "authors",
            localField: "authorID",
            foreignField : "_id",
            as : "authors"
        }
    },
    {
        $lookup : {
            from: "publishers",
            localField: "publisherID",
            foreignField : "_id",
            as : "publisher"
        }
    }
])
```
<img src="./screenshots/5.jpg" width="600">

6. Tampilkan summary data authors, books, dan publishers sesuai dengan Output.
```javascript
db.authors.aggregate([
    {
        $lookup : {
            from: "books",
            localField: "_id",
            foreignField : "authorID",
            as : "books"
        }
    },
    /* {
        $unwind: "$books"
    }, */
    {
        $lookup : {
            from: "publishers",
            localField: "books.publisherID",
            foreignField : "_id",
            as : "publishers"
        }
    },
    /* {
        $unwind: "$publishers"
    }, */
    {
        $project: {
            _id : { $concat: ["$firstName"," ","$lastName"]},
            number_of_books : {$size: "$books.authorID"},
            list_of_book : "$books.title",
        }
    }
])
```
<img src="./screenshots/6.jpg" width="600">

7. Digital_outlet ingin memberikan diskon untuk setiap buku, diskon di tentukan melihat harga buku tersebut dengan pembagian seperti ini.
```javascript
db.books.aggregate(
    [
       {
          $project:
            {
                _id : 1,
                title: 1,
                price: 1,
                discount:
                {
                    $cond: { if: { $lte: [ "$price", 60000] }, then: "1%", else: 
                    {
                        $cond: { if: { $lte : [ "$price", 90000] }, then: "2%", else: "3%" }
                    } }
                }
            }
       }
    ]
 )
```
<img src="./screenshots/7.jpg" width="600">

8. Tampilkan semua nama buku, harga, nama author dan nama publisher, urutkan dari harga termahal ke termurah.
```javascript
db.books.aggregate([
    {
        $sort: {price:-1}
    },
    {
        $lookup : {
            from: "authors",
            localField: "authorID",
            foreignField : "_id",
            as : "authors"
        }
    },
    {
        $unwind: "$authors"
    },
    {
        $lookup : {
            from: "publishers",
            localField: "publisherID",
            foreignField : "_id",
            as : "publishers"
        }
    },
    {
        $unwind: "$publishers"
    },
    {
        $project:
          {
              _id: 0,
              title: 1,
              price: 1,
              author : { $concat: ["$authors.firstName"," ", "$authors.lastName"]},
              publisher : "$publishers.publisherName"
          }
     }
])
```
<img src="./screenshots/8.jpg" width="600">

9. Tampilkan data nama buku harga dan publisher, kemudian tampilkan hanya data ke 3 dan ke 4.
```javascript
db.books.aggregate([
    {
        $lookup : {
            from: "publishers",
            localField: "publisherID",
            foreignField : "_id",
            as : "publishers"
        }
    },
    {
        $unwind: "$publishers"
    },
    {
        $project:
          {
              _id: 1,
              title: 1,
              price: 1,
              publisher : "$publishers.publisherName"
          }
     }
])
```
<img src="./screenshots/9a.jpg" width="600">

```javascript
db.books.aggregate([
    {
        $skip:2
    },
    {
        $limit:2
    },
    {
        $lookup : {
            from: "publishers",
            localField: "publisherID",
            foreignField : "_id",
            as : "publishers"
        }
    },
    {
        $unwind: "$publishers"
    },
    {
        $project:
          {
              _id: 1,
              title: 1,
              price: 1,
              publisher : "$publishers.publisherName"
          }
     }
])
```
<img src="./screenshots/9b.jpg" width="600">