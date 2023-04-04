package postgre

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetById(id int) (model.UserCreateResp, error) {
	var userResp model.UserCreateResp
	err := r.db.Table("users").First(&userResp, id)
	if err.Error != nil {
		return model.UserCreateResp{}, err.Error
	}
	return userResp, nil
}

func (r *UserRepository) Delete(id int) error {
	var user model.UserCreateReq
	result := r.db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = r.db.Delete(&user)
	return result.Error
}

func (r *UserRepository) Update(id int, user model.UserCreateReq) error {
	result := r.db.Where("id = ?", id).Updates(user)
	return result.Error
}
