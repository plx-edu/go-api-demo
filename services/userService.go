package services

import (
	"go-api-demo/db"
	"go-api-demo/models"
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