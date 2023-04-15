package service

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository"
)

type UserService struct {
	Repo *repository.Repository
}

func (s *UserService) TakeBook(id uint, bookId uint) error {
	return s.Repo.User.TakeBook(id, bookId)
}

func (s *UserService) ReturnBook(id uint, bookId uint) error {
	return s.Repo.User.ReturnBook(id, bookId)
}

func (s *UserService) GetUserBooks(id uint) ([]model.Book, error) {
	return s.Repo.User.GetUserBooks(id)
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

func (s *UserService) Update(id int, user model.User) error {
	return s.Repo.User.Update(id, user)
}
