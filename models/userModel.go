package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username	string	`gorm:"not null;default:null;uniqueIndex" json:"username"`
	Password	string	`gorm:"not null;default:null" json:"password"`
	Salt			string	`gorm:"not null;default:null" json:"salt"`
	// Todo			Todo		`gorm:"constraint:OnDelete:CASCADE;"`
}