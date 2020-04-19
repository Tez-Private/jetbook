package book

import (
	"context"
	"log"

	"github.com/Tez-Private/jetbook/db"

	model "github.com/Tez-Private/jetbook/model"
)

// Create is INSERT Book
func Create(ctx context.Context, data *GoogleBooksAPI) (*model.Books, error) {
	book := &model.Books{
		Title: data.Items[0].VolumeInfo.Title,
		Isbn:  data.Items[0].VolumeInfo.IndustryIdentifiers[0].Identifier,
	}

	err := db.Mysql.Create(&book).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Mysql.Where("ID=?", book.ID).First(&book).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return book, err
}

//GetAll is SELET * FROM users all
func GetAll(ctx context.Context) ([]model.Books, error) {
	var books []model.Books

	err := db.Mysql.Find(&books).Error
	if err != nil {
		log.Println("Nothing")
		return books, err
	}
	return books, err
}

//Get is SELET * FROM books where isbn = x;
func Get(isbn string) (model.Books, error) {
	var book model.Books

	err := db.Mysql.Where("isbn=?", isbn).First(&book).Error
	if err != nil {
		log.Println("Nothing")
		return book, err
	}
	return book, err
}

// Delete is
func Delete(ctx context.Context, isbn string) error {
	var searchID model.Books
	var book model.Books

	//err := db.Mysql.Where("isbn=?", isbn).Delete(&book).Error
	err := db.Mysql.Where("isbn=?", isbn).First(&searchID).Error
	if err != nil {
		log.Println()
		return err
	}

	err = db.Mysql.Delete(&book, searchID.ID).Error
	if err != nil {
		log.Println()
		return err
	}
	log.Println(err)
	return nil
}
