package handler

import (
	"net/http"
	"onelab/internal/model"

	"github.com/labstack/echo/v4"
)

// GiveBook      godoc
// @Tags         Library
// @Summary      Create user borrow book
// @Description  Create user borrow book
// @Accept       json
// @Produce      json
// @Param        input body model.CreateBookReader  true  "Create book reader"
// @Success 201 {string} string "Thanks for picking up the book"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router        /library/give-book [post]
func (h Manager) GiveBook(c echo.Context) error {
	var bookReader model.CreateBookReader

	if err := c.Bind(&bookReader); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.srv.Library.GiveBook(c.Request().Context(), bookReader); err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "Thanks for picking up the book")
}

// ReturnBook    godoc
// @Tags         Library
// @Summary      Create user return book
// @Description  Create user return book
// @Accept       json
// @Produce      json
// @Param        input body model.BookReader  true  "Create book return"
// @Success 200 {string} string "Your book is returned"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router        /library/return-book [post]
func (h Manager) ReturnBook(c echo.Context) error {
	var bookReturn model.BookReader
	if err := c.Bind(&bookReturn); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err := h.srv.Library.ReturnBook(c.Request().Context(), bookReturn)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Your book is returned")
}

// List          godoc
// @Tags         Library
// @Summary      Lists of borrowed books with income
// @Description  Lists of borrowed books with income
// @Accept       json
// @Produce      json
// @Success 200 {object} model.BookRentSummary
// @Failure 500 {object} errorResponse
// @Router        /library/all [get]
func (h Manager) List(c echo.Context) error {
	list, err := h.srv.Library.List(c.Request().Context())
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, list)
}
