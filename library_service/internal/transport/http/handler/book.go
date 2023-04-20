package handler

import (
	"fmt"
	"net/http"
	"onelab/internal/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateBook    godoc
// @Tags         books
// @Summary      Create book
// @Description  Create book
// @Accept       json
// @Produce      json
// @Param 		 input body model.Book  true  "Create book"
// @Success 201 {string} string "Your book is created <id>"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router        /books/create [post]
func (h Manager) CreateBook(c echo.Context) error {
	var book model.Book

	if err := c.Bind(&book); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.srv.Book.Create(c.Request().Context(), book)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, fmt.Sprintf("Your book is created %d", id))
}

// GetBookByID   godoc
// @Tags         books
// @Summary      Get book by id
// @Description  Get book by id
// @Accept       json
// @Produce      json
// @Param 		 id path int true "Book ID"
// @Success 302 {object} model.Book
// @Failure 400,404 {object} errorResponse
// @Router        /books/{id} [get]
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

	return c.JSON(http.StatusFound, book)
}

// GetAllBook    godoc
// @Tags         books
// @Summary      Get all book
// @Description  Get all book
// @Accept       json
// @Produce      json
// @Success 200 {object} model.Book
// @Failure 400 {object} errorResponse
// @Router        /books/ [get]
func (h Manager) GetAllBook(c echo.Context) error {
	books, err := h.srv.Book.GetAll(c.Request().Context())
	if err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}
