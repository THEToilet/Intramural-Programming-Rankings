package main

import (
	"net/http"
)

// GET /showwiki?q=
func index(writer http.ResponseWriter, request *http.Request) {
	api()
}
