package models

import (
	"gorm.io/gorm"
)

type (
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func NewUserModel(conn *gorm.DB) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}
