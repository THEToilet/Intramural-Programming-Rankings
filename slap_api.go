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

// Returning user information as a string
func getUserScore() string {

	users := loadFile("user.txt")
	top := "username  AcceptedCount  AcceptedCountRank   RatedPointSum \n ---------------------------------------------\n"
	for _, user := range users {
		resp, err := http.Get("https://kenkoooo.com/atcoder/atcoder-api/v2/user_info?user=" + user)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		top += atCoderInfoParse(resp)
	}
	return top
}

func atCoderInfoParse(response *http.Response) string {

	body, err := ioutil.ReadAll(response.Body)
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
	result := (string(atCoderInfo.UserId) + "   " + string(atCoderInfo.AcceptedCount) + "  " + string(atCoderInfo.AcceptedCountRank) + "      " + string(atCoderInfo.RatedPointSum) + "\n")

	return result
}

func getApi(name string) *AtCoderInfo {

	respo, err := http.Get("https://kenkoooo.com/atcoder/atcoder-api/v2/user_info?user=" + name)

	if err != nil {
		fmt.Println(err)
	}
	defer respo.Body.Close()

	return getAtCoderInfoStruct(respo)
}

func getAtCoderHistoryStruct(response *http.Response) []*AtCoderHistory {

	body, err := ioutil.ReadAll(response.Body)
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

func getAtCoderInfoStruct(response *http.Response) *AtCoderInfo {

	body, err := ioutil.ReadAll(response.Body)
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
