package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"onelab/internal/model"

	"github.com/labstack/echo/v4"
)

func (h Manager) LoginUser(c echo.Context) error {
	var input model.Login

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.srv.User.Auth(c.Request().Context(), model.User{
		Email:        input.Email,
		PasswordHash: input.Password,
	}); err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	token, err := h.srv.Jwt.GenerateToken(input.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (h Manager) CreateUser(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return err
	}

	err := h.srv.User.Create(c.Request().Context(), user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusOK, "User successfuly created")
}

func (h Manager) GetUserByID(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid 'id' parametr")
	}

	user, err := h.srv.User.Get(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h Manager) UpdateUser(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid 'id' parametr")
	}

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user.ID = uint(id)
	err = h.srv.User.Update(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusOK, "User successfuly updated")
}
