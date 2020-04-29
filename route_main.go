package main

import (
	"fmt"
	"net/http"
)

// GET /showwiki?q=
func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, api())
}

func test(writer http.ResponseWriter, request *http.Request) {
	test_sql()
	users := getUserInfo()
	var result string
	for _, user := range users {
		//▸-▸-fmt.Println(user)
		result = string(user.Id) + "   " + string(user.Name) + "   " + string(user.AcceptedCount) + "   " + string(user.AcceptedCountRank) + "   " + string(user.RatedPointSum) + "   " + string(user.RatedPointSumRank) + "   " + string(user.CreatedTime) + "  " + string(user.UpdatedTime) + "\n"
	}
	fmt.Fprintf(writer, result)
}
