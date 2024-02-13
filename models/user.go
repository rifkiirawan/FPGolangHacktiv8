package models

import (
	"project-mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age          int           `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,range(8|99)~Your age must be above 8 and below 99"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json: "products"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json: "products"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json: "products"`
	// Created_at *time.Time `json:"created_at,omitempty"`
	// UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
