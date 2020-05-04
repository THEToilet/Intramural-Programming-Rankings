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
	generateHTML(writer, getAtCoderHistoryStruct(vals.Get("q")), "index", "navbar", "userpage")
}

// GET /rank
func rank(writer http.ResponseWriter, request *http.Request) {
//	fmt.Printf("#%v", ranking())
	generateHTML(writer, ranking(), "index", "navbar", "rate.rank")
}
