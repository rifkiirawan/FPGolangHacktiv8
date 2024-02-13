package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string `json:"title" form:"title" valid:"required~Title of your photo is required"`
	Caption   string `json:"caption" form:"caption"`
	Photo_URL string `json:"photo_url" form:"photo_url" valid:"required~URL of your photo is required"`
	UserID    uint
	User      *User
	// Created_at *time.Time `json:"created_at,omitempty"`
	// UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
