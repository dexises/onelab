package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"onelab/internal/model"

	"github.com/labstack/echo/v4"
)

// LoginUser     godoc
// @Tags         Auth
// @Summary      User login
// @Description  User login
// @Accept       json
// @Produce      json
// @Param        input body model.Login  true  "Login into account"
// @Success 200 {object} tokenMessage
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router        /users/login [post]
func (h Manager) LoginUser(c echo.Context) error {
	var input model.Login

	if err := c.Bind(&input); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.srv.User.Auth(c.Request().Context(), model.User{
		Email:        input.Email,
		PasswordHash: input.Password,
	}); err != nil {
		return newErrorResponse(c, http.StatusUnauthorized, "Wrong email or password")
	}

	token, err := h.JWTAuth.GenerateJWT(input.Email)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tokenMessage{Message: token})
}

// CreateUser    godoc
// @Tags         Auth
// @Summary      User create
// @Description  User create
// @Accept       json
// @Produce      json
// @Param        input body model.User  true  "Create account"
// @Success 201 {string} string "Your user is created <id>"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router        /users/create [post]
func (h Manager) CreateUser(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.srv.User.Create(c.Request().Context(), user)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, fmt.Sprintf("Your user is created %d", id))
}

// GetUserByID   godoc
// @Tags         User manipulation
// @Summary      Get user
// @Description  Get user by id
// @Accept       json
// @Produce      json
// @Param 		 id path int true "User ID"
// @Success 302 {object} model.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router        /users/{id} [get]
func (h Manager) GetUserByID(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user, err := h.srv.User.Get(c.Request().Context(), id)
	if errors.Is(err, model.ErrRecordNotFound) {
		return newErrorResponse(c, http.StatusNotFound, err.Error())
	} else if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusFound, user)
}

// UpdateUser    godoc
// @Tags         User manipulation
// @Summary      Update user password
// @Description  Update user password
// @Accept       json
// @Produce      json
// @Param 		 id path int true "User ID"
// @Success 200 {string} string "User successfully updated"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router        /users/{id} [put]
func (h Manager) UpdateUser(c echo.Context) error {
	var user model.Login
	if err := c.Bind(&user); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err := h.srv.User.Update(c.Request().Context(), user)
	if errors.Is(err, model.ErrRecordNotFound) {
		return newErrorResponse(c, http.StatusNotFound, err.Error())
	} else if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User successfully updated")
}
