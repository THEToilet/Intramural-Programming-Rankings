package main

import (
	"net/http"
	"time"
)

func main() {

	// handle static assets
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(config.Static))
	// Deleate Prefix
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// GET index
	mux.HandleFunc("/index", index)

	// GET test
	mux.HandleFunc("/test", test)

	mux.HandleFunc("/list", list)

	// starting up the server
	server := &http.Server{
		Addr:           "127.0.0.1:11180",
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
