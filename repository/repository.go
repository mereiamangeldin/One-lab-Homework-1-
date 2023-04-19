package repository

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/config"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository/postgre"
)

type IUserRepository interface {
	GetById(id uint) (model.UserCreateResp, error)
	Delete(id uint) error
	Update(id uint, user model.User) error
	UpdatePassword(id uint, pass model.ChangePassword) error
	GetUserBooks(id uint) ([]model.Book, error)
	GetBalance(id uint) (model.UserBalance, error)
	UpdateBalance(balance model.UserBalance) error
}
type IAuthorizationRepository interface {
	CreateUser(user model.User) (uint, error)
	GetUser(user model.AuthUser) (model.User, error)
}
type IBookRepository interface {
	GetBooks() ([]model.Book, error)
	CreateBook(book model.Book) (uint, error)
	UpdateBook(id uint, book model.Book) error
	GetBookById(id uint) (model.Book, error)
	DeleteBook(id uint) error
	GetPrice(id uint) (model.BookPrice, error)
}

type Repository struct {
	User IUserRepository
	Auth IAuthorizationRepository
	Book IBookRepository
}

func New(cfg *config.Config) (*Repository, error) {
	pgDB, err := postgre.Dial(cfg.PgURL)
	if err != nil {
		return nil, err
	}
	userRep := postgre.NewUserRepository(pgDB)
	auth := postgre.NewAuthorizationRepository(pgDB)
	book := postgre.NewBookRepository(pgDB)
	return &Repository{User: userRep, Auth: auth, Book: book}, nil
}
