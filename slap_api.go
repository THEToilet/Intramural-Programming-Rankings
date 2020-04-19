package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func api() {
	values := url.Values{}
	values.Add("user", "Toilet")
	resp, err := http.Get("https://kenkoooo.com/atcoder/atcoder-api/v2/user_info" + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	execute(resp)
	respone, err := http.Get("https://kenkoooo.com/atcoder/atcoder-api/v2/user_info" + "?" + "user=toitenu")

	defer respone.Body.Close()
	execute(respone)

}

type TrainInfomation struct {
    user_id                    string    `json:"@user_id"`
    accepted_count             string    `json:"@accepted_count"`
    accepted_count_rank        string    `json:"@accepted_cout_rank"`
    rated_point_sum            string    `json:"@rated_point_sum"`
    rated_point_sum_rank       string    `json:"@rated_point_sum_rank"`
}
func execute(response *http.Response) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}

//func preservation(name string){
//}
