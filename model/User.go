package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserBalance struct {
	UserID  uint    `json:"user_id"`
	Balance float64 `json:"balance"`
}

func (u *UserBalance) TableName() string {
	return "user_balance"
}

func (u *User) TableName() string {
	return "users"
}

type UserCreateResp struct {
	ID      uint
	Name    string
	Surname string
}

type AuthUser struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type ChangePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"'`
}
