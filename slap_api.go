package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AtCoderInfo struct {
	UserId            string      `json:"user_id"`
	AcceptedCount     json.Number `json:"accepted_count"`
	AcceptedCountRank json.Number `json:"accepted_count_rank"`
	RatedPointSum     json.Number `json:"rated_point_sum"`
	RatedPointSumRank json.Number `json:"rated_point_sum_rank"`
}

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

func getApi(name string) *AtCoderInfo {
	print(name)
	respo, err := http.Get("https://kenkoooo.com/atcoder/atcoder-api/v2/user_info?user=" + name)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer respo.Body.Close()
	
	return ee(respo)
}

func api() string {
	users := loadFile("user.txt")
	var re string = "username  AcceptedCount  AcceptedCountRank   RatedPointSum \n ---------------------------------------------\n"
	for _, user := range users {
		resp, err := http.Get("https://kenkoooo.com/atcoder/atcoder-api/v2/user_info?user=" + user)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer resp.Body.Close()

		re += execute(resp)
	}
	return re
	/*
		for _, user := range users {
			res, err := http.Get("https://atcoder.jp/users/" + user + "/history/json")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()
			done(res)
		}
	*/
}

func execute(response *http.Response) string {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	//	fmt.Println(string(body))

	var info string = string(body)
	// Unmarshal結果の格納先である構造体のポインターを取得
	atCoderInfo := new(AtCoderInfo)

	// JSON文字列をバイト列にキャスト
	jsonBytes := []byte(info)

	// xJapanにバイト列を格納する
	if err := json.Unmarshal(jsonBytes, atCoderInfo); err != nil {
		fmt.Println(err)
	}
	fmt.Println("UserId : " + atCoderInfo.UserId)
	fmt.Println("AcceptedCount : " + atCoderInfo.AcceptedCount)
	fmt.Println("AcceptedCountRank : " + atCoderInfo.AcceptedCountRank)
	fmt.Println("RatedPointSum : " + atCoderInfo.RatedPointSum)
	fmt.Println("RatedPointSumRank : " + atCoderInfo.RatedPointSumRank)
	result := (string(atCoderInfo.UserId) + "   " + string(atCoderInfo.AcceptedCount) + "  " + string(atCoderInfo.AcceptedCountRank) + "      " + string(atCoderInfo.RatedPointSum) + "\n")

	return result

}

func done(response *http.Response) {

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	//	fmt.Println(string(body))

	var info string = string(body)
	// Unmarshal結果の格納先である構造体のポインターを取得
	// atCoderHistory := new(AtCoderHistory)
	var atCoderHistories []*AtCoderHistory
	// JSON文字列をバイト列にキャスト
	// atCoderHistoriesにバイト列を格納する
	var aaa = json.Unmarshal([]byte(info), &atCoderHistories)
	if aaa != nil {
		fmt.Println(aaa)
		return
	}

	for _, history := range atCoderHistories {
		fmt.Println(history.IsRated)
		fmt.Println("Place : " + history.Place)
		fmt.Println("OldRating : " + history.OldRating)
		fmt.Println("NewRating : " + history.NewRating)
		fmt.Println("Performance : " + history.Performance)
		fmt.Println("InnerPerformance : " + history.InnerPerformance)
		fmt.Println("ContestScreenName : " + history.ContestScreenName)
		fmt.Println("ContestName : " + history.ContestName)
		fmt.Println("ContestNameEn : " + history.ContestNameEn)
		fmt.Println("EndTime : " + history.EndTime)
	}
}


func ee(response *http.Response) *AtCoderInfo {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	//	fmt.Println(string(body))

	var info string = string(body)
	// Unmarshal結果の格納先である構造体のポインターを取得
	atCoderInfo := new(AtCoderInfo)

	// JSON文字列をバイト列にキャスト
	jsonBytes := []byte(info)

	// xJapanにバイト列を格納する
	if err := json.Unmarshal(jsonBytes, atCoderInfo); err != nil {
		fmt.Println(err)
	}
	fmt.Println("UserId : " + atCoderInfo.UserId)
	fmt.Println("AcceptedCount : " + atCoderInfo.AcceptedCount)
	fmt.Println("AcceptedCountRank : " + atCoderInfo.AcceptedCountRank)
	fmt.Println("RatedPointSum : " + atCoderInfo.RatedPointSum)
	fmt.Println("RatedPointSumRank : " + atCoderInfo.RatedPointSumRank)

	return atCoderInfo

}

