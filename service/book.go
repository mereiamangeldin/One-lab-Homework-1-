package service

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository"
)

type BookService struct {
	Repo *repository.Repository
}

func (s *BookService) GetBooks() ([]model.Book, error) {
	return s.Repo.Book.GetBooks()
}

func (s *BookService) CreateBook(book model.Book) (uint, error) {
	return s.Repo.Book.CreateBook(book)
}

func (s *BookService) UpdateBook(id uint, book model.Book) error {
	return s.Repo.Book.UpdateBook(id, book)
}

func (s *BookService) GetBookById(id int) (model.Book, error) {
	return s.Repo.Book.GetBookById(id)
}

func (s *BookService) DeleteBook(id int) error {
	return s.Repo.Book.DeleteBook(id)
}

func NewBookService(repo *repository.Repository) *BookService {
	return &BookService{Repo: repo}
}
