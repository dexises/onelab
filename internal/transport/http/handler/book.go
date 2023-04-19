package handler

import (
	"fmt"
	"net/http"
	"onelab/internal/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Manager) CreateBook(c echo.Context) error {
	var book model.Book

	if err := c.Bind(&book); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.srv.Book.Create(c.Request().Context(), book)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Your book is created %d", id))
}

func (h Manager) GetBookByID(c echo.Context) error {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return newErrorResponse(c, http.StatusBadRequest, "Invalid 'id' parameter")
	}

	book, err := h.srv.Book.Get(c.Request().Context(), idInt)
	if err != nil {
		return newErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, book)
}

func (h Manager) GetAllBook(c echo.Context) error {
	books, err := h.srv.Book.GetAll(c.Request().Context())
	if err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}
