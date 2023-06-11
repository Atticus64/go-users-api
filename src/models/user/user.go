package userModel

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Name string
	Age int
	Language string
}

func (Model) TableName() string {
	return "users"
}


