db.books.insertMany([
    {_id: 1, title: "Wawasan Pancasila", authorID: 1, publisherID: 1, price: 71200, stats: {page: 324, weight: 300}, publishedAt: new Date("2018-10-01"), category: ["social", "politics"]},
    {_id: 2, title: "Mata Air Keteladanan", authorID: 1, publisherID: 2, price: 106250, stats: {page: 672, weight: 650}, publishedAt: new Date("2017-09-01"), category: ["social", "politics"]},
    {_id: 3, title: "Revolusi Pancasila", authorID: 1, publisherID: 1, price: 54400, stats: {page: 220, weight: 500}, publishedAt: new Date("2015-05-01"), category: ["social", "politics"]},
    {_id: 4, title: "Self Driving", authorID: 2, publisherID: 1, price: 58650, stats:{page: 286, weight: 300}, publishedAt: new Date("2018-05-01"), category: ["self-development"]},
    {_id: 5, title: "Self Disruption", authorID: 2, publisherID: 2, price: 83300, stats: {page: 400, weight: 800}, publishedAt: new Date("2018-05-01"), category: ["self-development"]}
    ])

db.authors.insertMany([
    {_id:1, firstName: "Yudi", lastName: "Latif"},
    {_id:2, firstName: "Rhenald", lastName: "Kasali"}
    ])

db.publishers.insertMany([
    {_id: 1, publisherName: "Expose"},
    {_id: 2, publisherName: "Mizan"}
    ])

// 1. Gabungkan (menampilkan) data buku dari author id 1 dan author id 2.
db.books.find({ $or: [{authorID:1},{authorID:2}]})

// 2. Tampilkan daftar buku dan harga author id 1.
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

// 3. Tampilan total jumlah halaman buku author id 2.
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

// 4. Tampilkan semua field books and authors terkait.
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

// 5. Tampilkan semua field books, authors, dan publishers terkait.
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

// 6. Tampilkan summary data authors, books, dan publishers sesuai dengan Output.
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

// 7. Digital_outlet ingin memberikan diskon untuk setiap buku, diskon di tentukan melihat harga buku tersebut
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

// 8. Tampilkan semua nama buku, harga, nama author dan nama publisher, urutkan dari harga termahal ke termurah.
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

// 9. Tampilkan data nama buku harga dan publisher, kemudian tampilkan hanya data ke 3 dan ke 4.
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