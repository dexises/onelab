package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"onelab/internal/model"
	"strconv"
)

func (h Manager) CreateBook(c echo.Context) error {
	var book model.Book

	if err := c.Bind(&book); err != nil {
		//return c.JSON(http.StatusBadRequest, err)
		return err
	}

	if err := h.srv.Book.Create(c.Request().Context(), book); err != nil {
		//return c.JSON(http.StatusBadGateway, err)
		return err
	}

	return c.JSON(http.StatusOK, "Your book is created")
}

func (h Manager) GetBookByID(c echo.Context) error {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid id parameter")
	}

	book, err := h.srv.Book.Get(c.Request().Context(), idInt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (h Manager) GetAllBook(c echo.Context) error {
	books, err := h.srv.Book.GetAll(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, books)
}
