package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Task		string	`gorm:"not null;default:null" json:"task"`
	IsDone	bool		`gorm:"not null;default:null" json:"isDone"`
	TodoID	string	`gorm:"not null;default:null" json:"todoID"`
	Todo    Todo		`gorm:"constraint:OnDelete:CASCADE;foreignKey:TodoID"`
}