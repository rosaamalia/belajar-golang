package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID		uint	`json:"user_id"`
	User		*User	`json:"user"`
	PhotoID		uint	`json:"photo_id"`
	Photo		*Photo	`json:"photo"`
	Message		string	`gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}