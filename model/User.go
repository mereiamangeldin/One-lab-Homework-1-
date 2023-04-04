package model

import "gorm.io/gorm"

type UserCreateReq struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Pass    string `json:"password"`
}

func (u *UserCreateReq) TableName() string {
	return "users"
}

type UserCreateResp struct {
	ID      uint
	Name    string
	Surname string
}
