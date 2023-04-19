package handler

import (
	"net/http"
	"onelab/internal/model"

	"github.com/labstack/echo/v4"
)

func (h Manager) GiveBook(c echo.Context) error {
	var bookReader model.CreateBookReader

	if err := c.Bind(&bookReader); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.srv.Library.GiveBook(c.Request().Context(), bookReader); err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Thanks for picking up the book")
}

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

func (h Manager) List(c echo.Context) error {
	list, err := h.srv.Library.List(c.Request().Context())
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, list)
}

func (h Manager) ListMonth(c echo.Context) error {
	listMonth, err := h.srv.Library.ListMonth(c.Request().Context())
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())

	}
	return c.JSON(http.StatusOK, listMonth)
}
