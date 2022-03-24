package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"stock/utils/response_utils"
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

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	assert.EqualValues(t, http.StatusOK, rec.Code)
	assert.False(t, response.Error)
	assert.Equal(t, response.Type, "success")
	assert.EqualValues(t, "pong", response.Message)
	assert.Nil(t, response.Data)
}
