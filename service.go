package echo

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sort"
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
	remoteAddr := r.Header.Get("X-Forwarded-For")

	if remoteAddr == "" {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Fatalln(err)
		}
		remoteAddr = host
	}

	fmt.Fprintln(w, remoteAddr)
}

// Headers returns all headers in the request
func Headers(w http.ResponseWriter, r *http.Request) {
	keys := make([]string, 0, len(r.Header))
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		vals := r.Header[k]
		sort.Strings(vals)
		for _, v := range vals {
			fmt.Fprintf(w, "%s: %s\n", k, v)
		}
	}
}
