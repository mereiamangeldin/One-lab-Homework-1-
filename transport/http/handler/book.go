package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"net/http"
	"strconv"
)

func (h *Manager) GetBooks(c echo.Context) error {
	books, err := h.srv.Book.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, books)
}

func (h *Manager) CreateBook(c echo.Context) error {
	var book model.Book
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	id, err := h.srv.Book.CreateBook(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Manager) UpdateBook(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var book model.Book
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res := h.srv.Book.UpdateBook(uint(id), book)
	if res != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Book updated successfully",
	})
}
func (h *Manager) DeleteBook(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = h.srv.Book.DeleteBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Book deleted successfully",
	})
}

func (h *Manager) GetBookById(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	book, err := h.srv.Book.GetBookById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "No such book",
		})
	}
	return c.JSON(http.StatusOK, book)
}
