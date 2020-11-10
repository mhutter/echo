package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mhutter/echo"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	port := "8000"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	s := http.Server{
		Addr:           ":" + port,
		Handler:        echo.NewService(),
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Listening to %s\n", s.Addr)
	s.ListenAndServe()
}
