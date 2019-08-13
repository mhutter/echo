package echo

import (
	"fmt"
	"net/http"
)

// NewService returns a servic which facilitates the echoing of various
// HTTP request attributes.
func NewService() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/ip", IP)
	h.HandleFunc("/headers", Headers)
	return h
}

// IP returns the IP of the requestor
func IP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.RemoteAddr)
}

// Headers returns all headers in the request
func Headers(w http.ResponseWriter, r *http.Request) {
	for k, values := range r.Header {
		for _, v := range values {
			fmt.Fprintf(w, "%s: %s\n", k, v)
		}
	}
}
