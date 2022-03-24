package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingControllerPing(t *testing.T) {
	pingController := pingController{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := pingController.Ping(c)
	assert.Nil(t, err)
	assert.EqualValues(t, "\"pong\"", strings.TrimSpace(rec.Body.String()))
	assert.EqualValues(t, http.StatusOK, rec.Code)
}
