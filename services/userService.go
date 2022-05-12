package services

import (
	"go-api-demo/db"
	"go-api-demo/models"
	"log"
)

func GetUsers() []models.User{
	var users []models.User

	// if db.Instance.Where("bad = ?", 8).Find(&users).Error != nil {
	db.Instance.Find(&users)
	// if db.Instance.Find(&users).Error != nil {
	// 	println("Boop")
	// }
	return users
}

func CreateUser(newUser models.User) models.User{
	user := newUser

	err := db.Instance.Create(&user).Error
	if err != nil {
		log.Println(":::",err)
	}else{
		log.Println("::: Created user")
		log.Println(":::",user)
	}

	return user
}