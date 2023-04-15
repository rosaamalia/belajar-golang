package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name			string	`gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaURL	string	`gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social media URL is required"`
	UserID			uint	`json:"user_id"`
	User			*User	`json:"user"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}