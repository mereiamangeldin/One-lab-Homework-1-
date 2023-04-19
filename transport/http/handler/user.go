package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"log"
	"net/http"
	"strconv"
)

func (h *Manager) UpdateUser(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var user model.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res := h.srv.User.Update(uint(id), user)
	if res != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "User updated successfully",
	})
}
func (h *Manager) UpdatePassword(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var pass model.ChangePassword
	if err := c.Bind(&pass); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res := h.srv.Auth.UpdatePassword(uint(id), pass)
	if res != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Password changed successfully",
	})
}

func (h *Manager) DeleteUser(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = h.srv.User.Delete(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "User deleted successfully",
	})
}

func (h *Manager) GetUserById(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	user, err := h.srv.User.GetById(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "No such user",
		})
	}
	return c.JSON(http.StatusOK, user)
}
func (h *Manager) GetBalance(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userBalance, err := h.srv.User.GetBalance(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "No such user",
		})
	}
	return c.JSON(http.StatusOK, userBalance)
}

func (h *Manager) GetUserBooks(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	books, err := h.srv.User.GetUserBooks(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, books)
}
func (h *Manager) RentBook(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	bookId_ := c.Param("book_id")
	bookId, err := strconv.Atoi(bookId_)
	if err != nil {
		return err
	}
	var transactionRequest model.TransactionRequest
	if err := c.Bind(&transactionRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	transactionRequest.BookID = uint(bookId)
	transactionRequest.UserID = uint(id)
	transactionRequest.Operation = "rent"
	err = h.srv.User.BuyBook(transactionRequest)
	if err != nil {
		log.Println(1, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, model.SuccessResponse{Message: "Book is successfully rented"})
}

func (h *Manager) ReturnBook(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	bookId_ := c.Param("book_id")
	bookId, err := strconv.Atoi(bookId_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = h.srv.User.ReturnBook(uint(id), uint(bookId))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, model.SuccessResponse{Message: "Book is successfully returned"})
}

func (h *Manager) BuyBook(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	bookId_ := c.Param("book_id")
	bookId, err := strconv.Atoi(bookId_)
	var transactionRequest model.TransactionRequest
	if err := c.Bind(&transactionRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	transactionRequest.BookID = uint(bookId)
	transactionRequest.UserID = uint(id)
	transactionRequest.Operation = "purchase"
	err = h.srv.User.BuyBook(transactionRequest)
	if err != nil {
		log.Println(1, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, model.SuccessResponse{Message: "Book is successfully purchased"})
}
