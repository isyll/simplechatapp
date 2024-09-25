package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name     	string	`json:"name" gorm:"unique"`
	UserID		int
	Owner	 	User	`json:"owner" gorm:"foreignKey:UserID;references:ID"`
	Users		[]User	`json:"users" gorm:"many2many:group_users"`
}
