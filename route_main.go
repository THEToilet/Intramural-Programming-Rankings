package main

import (
	"net/http"
	"fmt"
)

// GET /showwiki?q=
func index(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "index", "navbar", "wiki.content")
	fmt.Println("うんこ")
}

