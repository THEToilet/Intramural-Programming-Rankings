package main

import (
	"fmt"
	"net/http"
)

// GET index
func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, getUserScore())
}

// GET test
func test(writer http.ResponseWriter, request *http.Request) {
	storeInSql()
	result := getUserInfo()
	fmt.Fprintf(writer, result)
}

// GET list
func list(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, getAtCoderInfoStruct("Toilet"), "ranking")
}
