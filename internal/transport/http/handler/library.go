package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"onelab/internal/model"
	"time"
)

func (h Manager) GiveBook(c echo.Context) error {
	var bookReader model.CreateBookReader

	if err := c.Bind(&bookReader); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.srv.Library.GiveBook(c.Request().Context(), bookReader); err != nil {
		//return c.JSON(http.StatusBadGateway, err)
		return err
	}

	return c.JSON(http.StatusOK, "Your book loan is created")
}

func (h Manager) ReturnBook(c echo.Context) error {
	var bookReturn model.BookReader
	if err := c.Bind(&bookReturn); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	bookReturn.ReturnedAt = time.Now()

	err := h.srv.Library.ReturnBook(c.Request().Context(), bookReturn)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusOK, "Your book is returned")
}

func (h Manager) List(c echo.Context) error {
	list, err := h.srv.Library.List(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, list)
}

func (h Manager) ListMonth(c echo.Context) error {
	listMonth, err := h.srv.Library.ListMonth(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, listMonth)
}
