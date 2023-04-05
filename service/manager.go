package service

import (
	"errors"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository"
)

type IAllUsersService interface {
	Get() ([]model.UserCreateResp, error)
	Create(user model.UserCreateReq) (model.UserCreateResp, error)
}
type IUserService interface {
	GetById(id int) (model.UserCreateResp, error)
	Delete(id int) error
	Update(id int, user model.UserCreateReq) error
}

type Manager struct {
	User     IUserService
	AllUsers IAllUsersService
}

func NewManager(repo *repository.Repository) (*Manager, error) {
	userSrv := NewUserService(repo)
	allUsersSrv := NewAllUsersService(repo)
	if repo == nil { // молодец что проверил, но лучше это проверить на 1 строке 
		return nil, errors.New("No storage given")
	}
	return &Manager{User: userSrv, AllUsers: allUsersSrv}, nil
}
