package postgre

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"gorm.io/gorm"
)

type AuthorizationRepository struct {
	db *gorm.DB
}

func (r *AuthorizationRepository) GetUser(user model.AuthUser) (model.User, error) {
	var userResp model.User
	err := r.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&userResp)
	if err.Error != nil {
		return model.User{}, err.Error
	}
	return userResp, nil
}

func (r *AuthorizationRepository) CreateUser(user model.User) (uint, error) {
	err := r.db.Create(&user)
	if err.Error != nil {
		return 0, err.Error
	}
	return user.ID, nil
}

func NewAuthorizationRepository(db *gorm.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}
