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
	req.Header.Set("Accept", "foo/bar")
	req.Header.Add("Foo", "bar")
	req.Header.Add("Foo", "baz")
	res := httptest.NewRecorder()
	svc.ServeHTTP(res, req)
	assert.Equal(t, "Accept: foo/bar\nFoo: bar\nFoo: baz\n", res.Body.String())
}
