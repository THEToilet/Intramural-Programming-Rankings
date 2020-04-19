package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func API() {
	values := url.Values{}
	values.Add("user", "Toilet")
	resp, err := http.Get("https://kenkoooo.com/atcoder/atcoder-api/v2/user_info" + "?" + values.Encode())

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	execute(resp)
}

func execute(response *http.Response) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
