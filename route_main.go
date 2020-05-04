package main

import (
	"fmt"
	"net/http"
)

// GET index
func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, api())
}

// GET test
func test(writer http.ResponseWriter, request *http.Request) {
	test_sql()
	result := getUserInfo()
	fmt.Fprintf(writer, result)
	//generateHTML(writer, dirwalk("./resources"), "index", "navbar", "list")
}
