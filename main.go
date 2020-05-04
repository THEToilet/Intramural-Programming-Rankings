package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	// handle static assets
	mux := http.NewServeMux()

	//files := http.FileServer(http.Dir(config.Static))
	//	files := http.FileServer(http.Dir("/public"))
	// Deleate Prefix
	//	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// GET index
	//	mux.HandleFunc("/index", index)

	// GET test
	mux.HandleFunc("/test", test)

	mux.HandleFunc("/list", list)

	mux.HandleFunc("/user", user)

	mux.HandleFunc("/rank", rank)

	mux.HandleFunc("/css/", serveResource)
	mux.HandleFunc("/js/", serveResource)

	// starting up the server
	server := &http.Server{
		Addr:           ":11000",
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "text/javascript"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
