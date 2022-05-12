package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model

	Name		string	`gorm:"not null;default:null" json:"name"`
	UserID	string	`gorm:"not null;default:null" json:"userID"`
	User    User		`gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID"`
}