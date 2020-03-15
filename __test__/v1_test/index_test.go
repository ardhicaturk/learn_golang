package v1_test_test

import (
	v1 "github.com/ardhicaturk/learn_golang/handler/v1"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexRoutes(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	e := echo.New()
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	c := e.NewContext(req, rr)

	c.SetPath("/")

	if assert.NoError(t, v1.HandlerIndex(c)) {
		assert.Equal(t, http.StatusOK, rr.Code)
	}
}