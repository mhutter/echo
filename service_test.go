package echo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mhutter/echo"
	"github.com/stretchr/testify/assert"
)

func TestEchoService(t *testing.T) {
	t.Parallel()
	svc := echo.NewService()

	assert.HTTPSuccess(t, svc.ServeHTTP, "GET", "/ip", nil)

	req, err := http.NewRequest("GET", "/ip", nil)
	assert.Nil(t, err)
	req.RemoteAddr = "Mah IP"
	res := httptest.NewRecorder()
	svc.ServeHTTP(res, req)
	assert.Equal(t, "Mah IP\n", res.Body.String())
}
