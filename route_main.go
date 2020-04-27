package main

import (
	"fmt"
	"net/http"
)

// GET /showwiki?q=
func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, api())
}
