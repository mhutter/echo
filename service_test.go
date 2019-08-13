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
	t.Parallel()

	assert.HTTPSuccess(t, svc.ServeHTTP, "GET", "/ip", nil)

	req, err := http.NewRequest("GET", "/ip", nil)
	assert.Nil(t, err)
	req.RemoteAddr = "Mah IP"
	res := httptest.NewRecorder()
	svc.ServeHTTP(res, req)
	assert.Equal(t, "Mah IP\n", res.Body.String())
}

func TestEchoHeaders(t *testing.T) {
	t.Parallel()

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
