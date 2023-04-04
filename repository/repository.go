package repository

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/config"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository/postgre"
)

type IAllUsersRepository interface {
	Get() ([]model.UserCreateResp, error)
	Create(user model.UserCreateReq) (model.UserCreateResp, error)
}
type IUserRepository interface {
	GetById(id int) (model.UserCreateResp, error)
	Delete(id int) error
	Update(id int, user model.UserCreateReq) error
}

type Repository struct {
	User     IUserRepository
	AllUsers IAllUsersRepository
}

func New(cfg *config.Config) (*Repository, error) {
	pgDB, err := postgre.Dial(cfg.PgURL)
	if err != nil {
		return nil, err
	}
	userRep := postgre.NewUserRepository(pgDB)
	allUsersRep := postgre.NewAllUsersRepository(pgDB)
	return &Repository{User: userRep, AllUsers: allUsersRep}, nil
}
