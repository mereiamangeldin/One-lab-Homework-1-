package service

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository"
)

type AllUsersService struct {
	Repo *repository.Repository
}

func (s *AllUsersService) Get() ([]model.UserCreateResp, error) {
	return s.Repo.AllUsers.Get()
}

func (s *AllUsersService) Create(user model.UserCreateReq) (model.UserCreateResp, error) {
	return s.Repo.AllUsers.Create(user)
}

func NewAllUsersService(repo *repository.Repository) *AllUsersService {
	return &AllUsersService{Repo: repo}
}
