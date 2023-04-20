package handler

import "github.com/labstack/echo/v4"

type errorResponse struct {
	Message string `json:"message"`
}

type tokenMessage struct {
	Message string `json:"message"`
}

func newErrorResponse(c echo.Context, statusCode int, message string) *echo.HTTPError {
	return echo.NewHTTPError(statusCode, errorResponse{Message: message})
}
