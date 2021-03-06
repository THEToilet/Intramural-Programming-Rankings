package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GET "https://kenkoooo.com/atcoder/atcoder-api/v2/user_info?user=" + name"
type AtCoderInfo struct {
	UserId            string      `json:"user_id"`
	AcceptedCount     json.Number `json:"accepted_count"`
	AcceptedCountRank json.Number `json:"accepted_count_rank"`
	RatedPointSum     json.Number `json:"rated_point_sum"`
	RatedPointSumRank json.Number `json:"rated_point_sum_rank"`
}

// GET "https://atcoder.jp/users/" + user + "/history/json"
type AtCoderHistory struct {
	IsRated                     bool        `json:"IsRated"`
	Place                       json.Number `json:"Place"`
	OldRating                   json.Number `json:"OldRating"`
	NewRating                   json.Number `json:"NewRating"`
	Performance                 json.Number `json:"Performance"`
	InnerPerformance            json.Number `json:"InnerPerformance"`
	ContestScreenName           string      `json:"ContestScreenName"`
	ContestName                 string      `json:"ContestName"`
	ContestNameEn               string      `json:"ContestNameEn"`
	EndTime/*time.Time*/ string             `json:"EndTime"`
}

func getUsersatCoderInfo() []*AtCoderInfo {
	users := loadFile("user.txt")
	atcoderInfos := make([]*AtCoderInfo, 0, 100)
	for _, user := range users {
		atcoderInfos = append(atcoderInfos, getAtCoderInfoStruct(user))
	}
	return atcoderInfos
}

func getAtCoderInfoStruct(name string) *AtCoderInfo {

	res, err := http.Get("https://kenkoooo.com/atcoder/atcoder-api/v2/user_info?user=" + name)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var info string = string(body)
	// Unmarshal結果の格納先である構造体のポインターを取得
	atCoderInfo := new(AtCoderInfo)

	// JSON文字列をバイト列にキャスト
	jsonBytes := []byte(info)

	// atCoderInfoにバイト列を格納する
	if err := json.Unmarshal(jsonBytes, atCoderInfo); err != nil {
		fmt.Println(err)
	}

	return atCoderInfo
}
func getAtCoderHistoryStruct(name string) []*AtCoderHistory {

	res, err := http.Get("https://atcoder.jp/users/" + name + "/history/json")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var info string = string(body)
	// Unmarshal結果の格納先である構造体のポインターを取得
	var atCoderHistories []*AtCoderHistory
	// atCoderHistoriesにバイト列を格納する
	if err := json.Unmarshal([]byte(info), &atCoderHistories); err != nil {
		fmt.Println(err)
	}

	return atCoderHistories
}
