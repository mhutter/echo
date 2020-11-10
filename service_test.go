package echo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mhutter/echo"
	"github.com/stretchr/testify/assert"
)

var svc = echo.NewService()

func TestEchoIP(t *testing.T) {
	assert.HTTPSuccess(t, svc.ServeHTTP, "GET", "/ip", nil)

	req, err := http.NewRequest("GET", "/ip", nil)
	assert.Nil(t, err)
	req.RemoteAddr = "127.0.0.1:56789"
	res := httptest.NewRecorder()
	svc.ServeHTTP(res, req)
	assert.Equal(t, "127.0.0.1\n", res.Body.String())
}

func TestEchoIPForwarded(t *testing.T) {
	req, err := http.NewRequest("GET", "/ip", nil)
	assert.Nil(t, err)
	req.Header.Set("X-Forwarded-For", "::1")
	req.RemoteAddr = "127.0.0.1:56789"
	res := httptest.NewRecorder()
	svc.ServeHTTP(res, req)
	assert.Equal(t, "::1\n", res.Body.String())
}

func TestEchoHeaders(t *testing.T) {
	assert.HTTPSuccess(t, svc.ServeHTTP, "GET", "/headers", nil)

	req, err := http.NewRequest("GET", "/headers", nil)
	assert.Nil(t, err)
	req.Header = http.Header{
		"Foo":    {"a", "c", "b"},
		"Accept": {"foo/bar"},
	}
	res := httptest.NewRecorder()
	svc.ServeHTTP(res, req)
	expected := `Accept: foo/bar
Foo: a
Foo: b
Foo: c
`
	assert.Equal(t, expected, res.Body.String())
}
