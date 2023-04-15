package postgre

import (
	"errors"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) TakeBook(id uint, bookId uint) error {
	var book model.BookHistory
	err := r.db.Where("client_id = ? AND book_id = ?", id, bookId).First(&book)
	if err == nil {
		return errors.New("book is already taken")
	}
	takenBook := model.BookHistory{
		BookId:   bookId,
		ClientId: id,
		TakenAt:  time.Now(),
	}
	res := r.db.Create(&takenBook)
	return res.Error
}

func (r *UserRepository) ReturnBook(id uint, bookId uint) error {
	var book model.BookHistory
	err := r.db.Where("client_id = ? AND book_id = ?", id, bookId).First(&book)
	if err.Error != nil {
		return err.Error
	}
	book.ReturnedAt = time.Now()
	r.db.Save(&book)
	return nil
}

func (r *UserRepository) GetUserBooks(id uint) ([]model.Book, error) {
	var books []model.Book
	var ids []uint
	r.db.Raw("select book_id from book_history where client_id = ? and returned_at = ?", id, time.Time{}).Pluck("id", &ids)
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

func (r *UserRepository) GetById(id int) (model.UserCreateResp, error) {
	var userResp model.UserCreateResp
	err := r.db.Table("users").Where("deleted_at IS NULL").First(&userResp, id)
	if err.Error != nil {
		return model.UserCreateResp{}, err.Error
	}
	return userResp, nil
}

func (r *UserRepository) Delete(id int) error {
	var user model.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = r.db.Delete(&user)
	return result.Error
}

func (r *UserRepository) Update(id int, user model.User) error {
	result := r.db.Where("id = ?", id).Updates(user)
	return result.Error
}
