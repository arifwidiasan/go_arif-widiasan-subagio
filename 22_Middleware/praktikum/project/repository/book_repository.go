package repository

import (
	"fmt"
	conn "project/config"
	"project/model"
)

func CreateBooks(book model.Book) error {
	res := conn.DB.Create(&book)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func GetAllBooks() []model.Book {
	books := []model.Book{}
	conn.DB.Find(&books)

	return books
}

func GetOneByIdBook(id int) (book model.Book, err error) {
	res := conn.DB.Where("id = ?", id).Find(&book)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func DeleteByIdBook(id int) error {
	var books model.Book
	res := conn.DB.Unscoped().Delete(&books, id)

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func UpdateOneByIdBook(id int, book model.Book) error {
	res := conn.DB.Where("id = ?", id).UpdateColumns(&book)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}
