package service

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository"
)

type UserService struct {
	Repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetById(id int) (model.UserCreateResp, error) {
	return s.Repo.User.GetById(id)
}

func (s *UserService) Delete(id int) error {
	return s.Repo.User.Delete(id)
}

func (s *UserService) Update(id int, user model.UserCreateReq) error {
	return s.Repo.User.Update(id, user)
}
