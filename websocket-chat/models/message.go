package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID	int
	Sender     	User	`json:"sender" gorm:"foreignKey:SenderID;references:ID"`
	RcptID		int
	Rcpt 		User	`json:"rcpt" gorm:"foreignKey:RcptID;references:ID"`
	GroupID		int
	Group		Group 	`json:"group" gorm:"foreignKey:GroupID;references:ID"`
	Content		string	`json:"content"`
}
