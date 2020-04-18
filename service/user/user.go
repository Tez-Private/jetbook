package user

import (
	"context"
	"log"

	"github.com/Tez-Private/jetbook/db"
	model "github.com/Tez-Private/jetbook/model"
)

// Create is INSERT Users Table
func Create(ctx context.Context, params *CreateParams) (*model.Users, error) {
	user := &model.Users{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
	}
	err := db.Mysql.Create(&user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Mysql.Where("ID=?", user.ID).First(&user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, err
}

//GetAll is SELET * FROM users all
func GetAll(ctx context.Context) ([]model.Users, error) {
	var users []model.Users

	err := db.Mysql.Find(&users).Error
	if err != nil {
		log.Println("Nothing")
		return users, err
	}
	return users, err
}

//Get is SELET * FROM users where id = x;
func Get(id string) (model.Users, error) {
	var user model.Users

	err := db.Mysql.Where("ID=?", id).First(&user).Error
	if err != nil {
		log.Println("nothing")
		return user, err
	}
	return user, err
}

//Update is
func Update(ctx context.Context, params *UpdateParams, id string) (*model.Users, error) {
	user := &model.Users{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
	}

	err := db.Mysql.Table("users").Where("ID=?", id).Update(&user).Error
	if err != nil {
		return nil, err
	}

	err = db.Mysql.Where("ID=?", id).First(&user).Error
	if err != nil {
		log.Println("nothing")
		return user, err
	}

	return user, err
}

//Delete is
func Delete(ctx context.Context, id string) error {
	var user model.Users
	err := db.Mysql.Delete(&user, id).Error
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(err)
	return nil
}
