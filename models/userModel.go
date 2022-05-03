package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name string 
	Age  int
	// Name string `json "name"`
	// Age  int    `json "age"`
}