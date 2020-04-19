package rent

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Tez-Private/jetbook/db"
	model "github.com/Tez-Private/jetbook/model"
)

// Create is INSERT Rents Table
func Create(ctx context.Context, params *CreateParams) (*model.Rents, error) {
	rent := &model.Rents{
		BookID: params.BookID,
		UserID: params.UserID,
	}
	// After 2 weeks
	t := time.Now()
	expected := t.Add(336 * time.Hour)
	rent.ExpectedReturnDate = expected
	rent.BackCheck = 0

	book := &model.Books{}
	recordNotFound := db.Mysql.Table("books").Where("ID=? AND Rent=0", rent.BookID).First(&book).RecordNotFound()

	if recordNotFound == true {
		log.Println(recordNotFound)
		return nil, fmt.Errorf("This Book already has been rended")
	}

	err := db.Mysql.Table("books").Where("ID=?", rent.BookID).Update("rent", 1).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Mysql.Create(&rent).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Mysql.Where("ID=?", rent.ID).First(&rent).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rent, err
}

//GetAll is SELET * FROM rents all
func GetAll(ctx context.Context) ([]model.Rents, error) {
	var rents []model.Rents

	err := db.Mysql.Find(&rents).Error
	if err != nil {
		log.Println("Nothing")
		return rents, err
	}
	return rents, err
}

//Get is SELET * FROM rents where id = x;
func Get(id string) (model.Rents, error) {
	var rent model.Rents

	err := db.Mysql.Where("ID=?", id).First(&rent).Error
	if err != nil {
		log.Println("nothing")
		return rent, err
	}
	return rent, err
}

//Return is
func Return(ctx context.Context, params *UpdateParams, id string) (*model.Rents, error) {
	rent := &model.Rents{}

	err := db.Mysql.Table("rents").Where("ID=?", id).Update(&rent.BackCheck, 1).Error
	if err != nil {
		return nil, err
	}

	err = db.Mysql.Where("ID=?", id).First(&rent).Error
	if err != nil {
		log.Println("nothing")
		return rent, err
	}

	return rent, err
}

//Update is
func Update(ctx context.Context, params *UpdateParams, id string) (*model.Rents, error) {
	rent := &model.Rents{
		BookID:             params.BookID,
		UserID:             params.UserID,
		ExpectedReturnDate: params.ExpectedReturnDate,
		BackCheck:          params.BackCheck,
	}

	err := db.Mysql.Table("rents").Where("ID=?", id).Update(&rent).Error
	if err != nil {
		return nil, err
	}

	err = db.Mysql.Where("ID=?", id).First(&rent).Error
	if err != nil {
		log.Println("nothing")
		return rent, err
	}

	return rent, err
}

//Delete is
/*
func Delete(ctx context.Context, id string) error {
	var rent model.Rents
	err := db.Mysql.Delete(&rent, id).Error
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(err)
	return nil
}
*/
