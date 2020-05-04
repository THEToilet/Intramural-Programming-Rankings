package main

import (
	"fmt"
	"net/http"
)

/*
// GET index
func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, getUserScore())
}
*/
// GET test
func test(writer http.ResponseWriter, request *http.Request) {
	storeInSql()
	result := getUserInfo()
	fmt.Fprintf(writer, result)
}

// GET list
func list(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, getUsersatCoderInfo(), "index", "navbar", "ranking")
}

// GET /user?q=
func user(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	fmt.Println(vals.Get("q"))
	generateHTML(writer, getAtCoderHistoryStruct(vals.Get("q")), "index", "navbar", "userpage")
}
