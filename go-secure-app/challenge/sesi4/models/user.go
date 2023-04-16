package models

import (
	"sesi4/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username		string			`gorm:"uniqueIndex; not null" json:"username" form:"username" valid:"required~Username is required"`
	Email			string			`gorm:"uniqueIndex; not null" json:"email" form:"email" valid:"required~Email is required, email~Invalid email format"`
	Password		string			`gorm:"not null" json:"password" form:"password" valid:"required~Password is required, minstringlength(6)~Password has to have minimum length of 6 characters"`
	Age				int				`gorm:"not null" json:"age" form:"age" valid:"required~Age is required, numeric~Invalid age format. Age must be number, range(8|100)~Minimum age is 8 years old"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	// hash password sebelum disimpan
	u.Password = helpers.HashPass(u.Password)
	
	err = nil
	return
}