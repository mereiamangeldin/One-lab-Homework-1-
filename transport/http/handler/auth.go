package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"net/http"
)

// SignIn
// @Summary SignIn
// @Description  registration process
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input    body model.User  true  "account info"
// @Router       auth/sign-in

func (h *Manager) SignIn(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	id, err := h.srv.Auth.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]uint{
		"id": id,
	})
}

// SignUp
// @Summary SignUp
// @Description  authorization process
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input    body model.AuthUser  true  "account info"
// @Router       auth/sign-up

func (h *Manager) SignUp(c echo.Context) error {
	var input model.AuthUser
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	token, err := h.srv.Auth.GenerateToken(input)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{
			"error": "invalid password or username",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
