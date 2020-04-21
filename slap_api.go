package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
	"time"
)

type AtCoderInfo struct {
	UserId              string `json:"user_id"`
	AcceptedCount       int64  `json:"accepted_count"`
	AcceptedCountRank   int64  `json:"accepted_cout_rank"`
	RatedPointSum       int64  `json:"rated_point_sum"`
	RatedPointSumRank   int64  `json:"rated_point_sum_rank"`
}

type AtCoderInfo []AtCoderInfo

type AtCoderHistory struct {
	IsRated             bool      `json:"IsRated"`
	Place               int64     `json:"Place"`
	OldRating           int64     `json:"OldRating"`
	NewRating           int64     `json:"NewRating"`
	Perfomance          int64     `json:"Perfomance"`
	InnerPerformance    int64     `json:"InnerPerformance"`
	ContestScreenName   string    `json:"ContestScreenName"`
	ContestName         string    `json:"ContestName"`
	ContestNameEn       string    `json:"ContestNameEn"`
	EndTime             time.Time `json:"EndTime"`
}

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
	
  respones, err := http.Get("https://kenkoooo.com/atcoder/atcoder-api/v2/user_info" + "?" + "user=5jiKinoko")
	defer respones.Body.Close()
	execute(respones)
  
  responese, err := http.Get("https://atcoder.jp/users/Toilet/history/json")
	defer responese.Body.Close()
	execute(responese)
}

func execute(response *http.Response) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}
