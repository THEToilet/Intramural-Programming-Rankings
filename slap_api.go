package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type AtCoderInfo struct {
	UserId            string      `json:"user_id"`
	AcceptedCount     json.Number `json:"accepted_count"`
	AcceptedCountRank json.Number `json:"accepted_count_rank"`
	RatedPointSum     json.Number `json:"rated_point_sum"`
	RatedPointSumRank json.Number `json:"rated_point_sum_rank"`
}

//type AtCoderInfo []AtCoderInfo

type AtCoderHistory struct {
	IsRated           bool      `json:"IsRated"`
	Place             int       `json:"Place"`
	OldRating         int       `json:"OldRating"`
	NewRating         int       `json:"NewRating"`
	Perfomance        int       `json:"Perfomance"`
	InnerPerformance  int       `json:"InnerPerformance"`
	ContestScreenName string    `json:"ContestScreenName"`
	ContestName       string    `json:"ContestName"`
	ContestNameEn     string    `json:"ContestNameEn"`
	EndTime           time.Time `json:"EndTime"`
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

	//	responese, err := http.Get("https://atcoder.jp/users/Toilet/history/json")
	//	defer responese.Body.Close()
	//	execute(responese)
}

func execute(response *http.Response) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))

	var info string = string(body)
	// Unmarshal結果の格納先である構造体のポインターを取得
	atCoderInfo := new(AtCoderInfo)

	// JSON文字列をバイト列にキャスト
	jsonBytes := []byte(info)

	// xJapanにバイト列を格納する
	if err := json.Unmarshal(jsonBytes, atCoderInfo); err != nil {
		fmt.Println(err)
		return
	}
	// for _, members := range xJapan.Members {
	//	fmt.Printf("UserId: %-7s AcceptedCount: %d AcceptedCountRank: %d RatedPointSum: %d RatedPointSumRank: %d\n", atCoderInfo.AcceptedCount, atCoderInfo.AcceptedCountRank,
	//	atCoderInfo.RatedPointSum, atCoderInfo.RatedPointSumRank)
	fmt.Println("UserId : " + atCoderInfo.UserId)
	fmt.Println("AcceptedCount : " + atCoderInfo.AcceptedCount)
	fmt.Println("AcceptedCountRank : " + atCoderInfo.AcceptedCountRank)
	fmt.Println("RatedPointSum : " + atCoderInfo.RatedPointSum)
	fmt.Println("RatedPointSumRank : " + atCoderInfo.RatedPointSumRank)
	// }
}
