package service

import (
	"errors"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository"
)

type IUserService interface {
	GetById(id int) (model.UserCreateResp, error)
	Delete(id int) error
	Update(id int, user model.User) error
	TakeBook(id uint, bookId uint) error
	ReturnBook(id uint, bookId uint) error
	GetUserBooks(id uint) ([]model.Book, error)
}
type IAuthorizationService interface {
	CreateUser(user model.User) (uint, error)
	GenerateToken(user model.AuthUser) (string, error)
	ParseToken(token string) (uint, error)
	UpdatePassword(id uint, pass model.ChangePassword) error
}

type IBookService interface {
	GetBooks() ([]model.Book, error)
	CreateBook(book model.Book) (uint, error)
	UpdateBook(id uint, book model.Book) error
	GetBookById(id int) (model.Book, error)
	DeleteBook(id int) error
}

type Manager struct {
	User IUserService
	Auth IAuthorizationService
	Book IBookService
}

func NewManager(repo *repository.Repository) (*Manager, error) {
	userSrv := NewUserService(repo)
	authSrv := NewAuthorizationService(repo)
	bookSrv := NewBookService(repo)
	if repo == nil {
		return nil, errors.New("No storage given")
	}
	return &Manager{User: userSrv, Auth: authSrv, Book: bookSrv}, nil
}
