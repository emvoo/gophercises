package models

import (
	"github.com/jinzhu/gorm"
)

type Phone struct {
	gorm.Model
	Number string `gorm:"type:varchar(15)"`
}

func (Phone) TableName() string {
	return "phone_numbers"
}
