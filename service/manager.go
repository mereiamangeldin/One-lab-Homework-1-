package service

import (
	"errors"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository"
)

type IUserService interface {
	GetById(id uint) (model.UserCreateResp, error)
	Delete(id uint) error
	Update(id uint, user model.User) error
	BuyBook(transactionReq model.TransactionRequest) error
	ReturnBook(id uint, bookId uint) error
	GetUserBooks(id uint) ([]model.Book, error)
	GetBalance(id uint) (model.UserBalance, error)
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
	GetBookById(id uint) (model.Book, error)
	DeleteBook(id uint) error
}

type Manager struct {
	User IUserService
	Auth IAuthorizationService
	Book IBookService
}

func NewManager(repo *repository.Repository) (*Manager, error) {
	if repo == nil {
		return nil, errors.New("No storage given")
	}
	userSrv := NewUserService(repo)
	authSrv := NewAuthorizationService(repo)
	bookSrv := NewBookService(repo)
	return &Manager{User: userSrv, Auth: authSrv, Book: bookSrv}, nil
}
