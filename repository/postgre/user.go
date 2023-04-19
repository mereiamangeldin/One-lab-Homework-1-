package postgre

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) UpdateBalance(balance model.UserBalance) error {
	result := r.db.Where("user_id = ?", balance.UserID).Updates(balance)
	return result.Error
}

func (r *UserRepository) GetBalance(id uint) (model.UserBalance, error) {
	var balance model.UserBalance
	err := r.db.Where("user_id = ?", id).First(&balance)
	if err.Error != nil {
		return model.UserBalance{}, err.Error
	}
	return balance, nil
}

func (r *UserRepository) GetUserBooks(id uint) ([]model.Book, error) {
	var books []model.Book
	var ids []uint
	r.db.Raw("select book_id from transactions where client_id = ? and returned_at = ?", id, time.Time{}).Pluck("id", &ids)
	err := r.db.Where("id IN (?)", ids).Find(&books)
	if err.Error != nil {
		return nil, err.Error
	}

	return books, nil
}

func (r *UserRepository) UpdatePassword(id uint, pass model.ChangePassword) error {
	var user model.User
	err := r.db.Where("password = ?", pass.CurrentPassword).First(&user)
	if err.Error != nil {
		return err.Error
	}
	user.Password = pass.NewPassword
	r.db.Save(&user)
	return nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetById(id uint) (model.UserCreateResp, error) {
	var userResp model.UserCreateResp
	err := r.db.Table("users").Where("deleted_at IS NULL").First(&userResp, id)
	if err.Error != nil {
		return model.UserCreateResp{}, err.Error
	}
	return userResp, nil
}

func (r *UserRepository) Delete(id uint) error {
	var user model.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = r.db.Delete(&user)
	return result.Error
}

func (r *UserRepository) Update(id uint, user model.User) error {
	result := r.db.Where("id = ?", id).Updates(user)
	return result.Error
}
