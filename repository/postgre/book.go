package postgre

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func (r *BookRepository) GetBooks() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Find(&books)
	if err.Error != nil {
		return nil, err.Error
	}
	return books, nil
}

func (r *BookRepository) CreateBook(book model.Book) (uint, error) {
	res := r.db.Create(&book)
	return book.Id, res.Error
}

func (r *BookRepository) UpdateBook(id uint, bookReq model.Book) error {
	var book model.Book
	result := r.db.Model(&book).Where("id = ?", id).Updates(bookReq)
	return result.Error
}

func (r *BookRepository) GetBookById(id int) (model.Book, error) {
	var book model.Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return model.Book{}, result.Error
	}
	return book, nil
}

func (r *BookRepository) DeleteBook(id int) error {
	var book model.Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return result.Error
	}
	result = r.db.Delete(&book)
	return result.Error
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}
