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
	return h
}

// IP returns the IP of the requestor
func IP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.RemoteAddr)
}
