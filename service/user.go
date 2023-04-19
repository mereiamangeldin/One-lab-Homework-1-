package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository"
	"net/http"
)

const transactionUrl = "http://transaction:8000/transactions"

type UserService struct {
	Repo *repository.Repository
}

func (s *UserService) GetBalance(id uint) (model.UserBalance, error) {
	return s.Repo.User.GetBalance(id)
}

func (s *UserService) BuyBook(transactionReq model.TransactionRequest) error {
	_, err := s.Repo.Book.GetBookById(transactionReq.BookID)
	if err != nil {
		return errors.New("no such book")
	}
	_, err = s.Repo.User.GetById(transactionReq.UserID)
	if err != nil {
		return errors.New("no such user")
	}
	bookPrice, err := s.Repo.Book.GetPrice(transactionReq.BookID)
	if err != nil {
		return err
	}
	userBalance, err := s.Repo.User.GetBalance(transactionReq.UserID)
	if err != nil {
		return err
	}
	var RealPrice float64
	if transactionReq.Operation == "rent" {
		RealPrice = bookPrice.RentalPrice
	} else if transactionReq.Operation == "purchase" {
		RealPrice = bookPrice.PurchasePrice
	}
	if userBalance.Balance < bookPrice.RentalPrice {
		return errors.New("not enough funds")
	}
	transaction := model.Transaction{
		BookID:          transactionReq.BookID,
		ClientID:        transactionReq.UserID,
		Amount:          RealPrice,
		TransactionType: transactionReq.Operation,
	}
	err = s.createTransaction(transaction)
	if err != nil {
		return err
	}
	userBalance.Balance -= bookPrice.RentalPrice
	err = s.Repo.User.UpdateBalance(userBalance)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserService) createTransaction(transaction model.Transaction) error {
	transactionJson, err := json.Marshal(transaction)
	if err != nil {
		return err
	}
	resp, err := http.Post(transactionUrl, "application/json", bytes.NewBuffer(transactionJson))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("remote service returned error: %s", resp.Status)
	}
	return nil
}

func (s *UserService) ReturnBook(id uint, bookId uint) error {
	_, err := s.Repo.Book.GetBookById(bookId)
	if err != nil {
		return errors.New("no such book")
	}
	_, err = s.Repo.User.GetById(id)
	if err != nil {
		return errors.New("no such user")
	}

	return nil
}

func (s *UserService) GetUserBooks(id uint) ([]model.Book, error) {
	return s.Repo.User.GetUserBooks(id)
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetById(id uint) (model.UserCreateResp, error) {
	return s.Repo.User.GetById(id)
}

func (s *UserService) Delete(id uint) error {
	return s.Repo.User.Delete(id)
}

func (s *UserService) Update(id uint, user model.User) error {
	return s.Repo.User.Update(id, user)
}
