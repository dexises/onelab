package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Manager) CreateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	// var user model.User

	// if err := c.
}

func (h Manager) GetUserByID(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello world")
}

func (h Manager) UpdateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello world")
}
