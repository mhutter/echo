package echo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
	"github.com/mhutter/echo"
)

var svc = echo.NewService()

func TestEchoIP(t *testing.T) {
	is := is.New(t)
	req := httptest.NewRequest("GET", "/ip", nil)
	req.RemoteAddr = "127.0.0.1:56789"
	res := httptest.NewRecorder()

	svc.ServeHTTP(res, req)

	is.Equal("127.0.0.1\n", res.Body.String())
}

func TestEchoIPForwarded(t *testing.T) {
	is := is.New(t)
	req := httptest.NewRequest("GET", "/ip", nil)
	req.Header.Set("X-Forwarded-For", "::1")
	req.RemoteAddr = "127.0.0.1:56789"
	res := httptest.NewRecorder()

	svc.ServeHTTP(res, req)

	is.Equal("::1\n", res.Body.String())
}

func TestEchoHeaders(t *testing.T) {
	is := is.New(t)
	req := httptest.NewRequest("GET", "/headers", nil)
	res := httptest.NewRecorder()

	svc.ServeHTTP(res, req)

	is.Equal(res.Code, 200)

	req = httptest.NewRequest("GET", "/headers", nil)
	req.Header = http.Header{
		"Foo":    {"a", "c", "b"},
		"Accept": {"foo/bar"},
	}
	res = httptest.NewRecorder()

	svc.ServeHTTP(res, req)

	expected := `Accept: foo/bar
Foo: a
Foo: b
Foo: c
`
	is.Equal(res.Body.String(), expected)
}
