package handler

import (
	"net/http"
	"onelab/internal/model"

	"github.com/labstack/echo/v4"
)

func (h Manager) TransactionsCreate(c echo.Context) error {
	var transaction model.TransactionCreate

	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, model.TransactionResponse{
			ID:      0,
			Success: false,
			Message: err.Error(),
		})
	}

	id, err := h.srv.Transactions.CreateTransaction(c.Request().Context(), transaction)
	if err != nil {
		return c.JSON(http.StatusBadGateway, model.TransactionResponse{
			ID:      0,
			Success: false,
			Message: err.Error(),
		})
	}

	resp := &model.TransactionResponse{
		ID:      id,
		Success: true,
		Message: "",
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) AllTransactions(c echo.Context) error {
	transactions, err := h.srv.Transactions.GetAll(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transactions)
}
