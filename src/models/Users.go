package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirsName string `gorm:"not mull"`
	LastName string `gorm:"not mull"`
	Email    string `gorm:"not mull; unique_index"`
	Task 	 []Task
}