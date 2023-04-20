package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"onelab/internal/config"
	"onelab/internal/repository"
	"onelab/internal/service"
	"strings"
	"testing"
)

var (
	userJSON = `{"name":"Jon Snow1","email":"jon2@labstack.com","password":"123password"}`
)

func TestManager_CreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := createUserHandler(t)

	// Assertions
	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func createUserHandler(t *testing.T) *Manager {
	cfg := config.New()
	repo, err := repository.NewRepository(context.Background(), cfg)
	if err != nil {
		t.Errorf(err.Error())
	}
	srvManager := service.NewService(cfg, repo)
	return NewManager(cfg, srvManager)
}

//func TestManager_GetUser(t *testing.T) {
//	// Setup
//	e := echo.New()
//	req := httptest.NewRequest(http.MethodGet, "/", nil)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	c.SetPath("/users/:email")
//	c.SetParamNames("email")
//	c.SetParamValues("jon@labstack.com")
//	h := &handler{mockDB}
//
//	// Assertions
//	if assert.NoError(t, h.getUser(c)) {
//		assert.Equal(t, http.StatusOK, rec.Code)
//		assert.Equal(t, userJSON, rec.Body.String())
//	}
//}
