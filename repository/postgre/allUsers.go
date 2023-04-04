package postgre

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"gorm.io/gorm"
)

type AllUsersRepository struct {
	db *gorm.DB
}

func NewAllUsersRepository(db *gorm.DB) *AllUsersRepository {
	return &AllUsersRepository{db: db}
}

func (r *AllUsersRepository) Get() ([]model.UserCreateResp, error) {
	var users []model.UserCreateResp
	err := r.db.Table("users").Where("deleted_at is NULL").Find(&users)
	if err.Error != nil {
		return nil, err.Error
	}
	return users, nil
}

func (r *AllUsersRepository) Create(user model.UserCreateReq) (model.UserCreateResp, error) {
	var userResp model.UserCreateResp
	err := r.db.Create(&user)
	if err.Error != nil {
		return model.UserCreateResp{}, err.Error
	}
	//fmt.Println(user.ID)
	err = r.db.Table("users").First(&userResp, user.ID)
	if err.Error != nil {
		return model.UserCreateResp{}, err.Error
	}
	return userResp, nil
}
