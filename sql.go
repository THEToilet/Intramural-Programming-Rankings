package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

func parseJson(response *http.Response) *AtCoderInfo {
	bod, er := ioutil.ReadAll(response.Body)
	if er != nil {
		panic(er)
	}
	//▸-fmt.Println(string(body))

	var info string = string(bod)
	// Unmarshal結果の格納先である構造体のポインターを取得
	aCoderInfo := new(AtCoderInfo)

	// JSON文字列をバイト列にキャスト
	jsonBytes := []byte(info)

	// xJapanにバイト列を格納する
	if err := json.Unmarshal(jsonBytes, aCoderInfo); err != nil {
		fmt.Println(err)
	}
	return aCoderInfo
}

func test_sql() {
	// db接続
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	uusers := loadFile("user.txt")

	//userEx := Users{}

	uuuu := []Users{}
	for _, user := range uusers {
	  tmpInfo := getApi(user)
		print(tmpInfo)
		e := db.Where("Name = ?", user).Find(&uuuu)
		if e != nil {
			
			fmt.Println("#%v",tmpInfo)
			db.Create(&Users{
				Name:              user,
				AcceptedCount:     4,//int(tmpInfo),
				AcceptedCountRank: 4,//int(tmpInfo),
				RatedPointSum:     4,//int(tmpInfo),
				RatedPointSumRank: 4,//int(tmpInfo),
				CreatedTime:       getDate(),
				UpdatedTime:       getDate(),
			})
		}
	}

	//	fmt.Printf("%#v\n",db)

	/*
			   	userEx.Name = "Toilet"
			   	userEx.AcceptedCount = 123
			   	userEx.AcceptedCountRank = 1000000
			   	userEx.RatedPointSum = 34
			   	userEx.RatedPointSumRank = 12
			   	userEx.CreatedTime = getDate()
			   	userEx.UpdatedTime = getDate()
		   	// INSERTを実行
			     // db.Create(&userEx)
			     db.Update(&userEx)
	*/
	error := db.Create(&Users{
		Name:              "Unko",
		AcceptedCount:     123,
		AcceptedCountRank: 12,
		RatedPointSum:     34,
		RatedPointSumRank: 12,
		CreatedTime:       getDate(),
		UpdatedTime:       getDate(),
	}).Error

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("データ追加成功")
	}

	//データを格納する変数を定義
	users := []Users{}

	//全取得
	db.Find(&users)
	//	print(users)
	//表示
	for _, user := range users {
		//		fmt.Println(user)
		fmt.Println(user.Id)
		fmt.Println(user.Name)
		fmt.Println(user.AcceptedCount)
		fmt.Println(user.AcceptedCountRank)
		fmt.Println(user.RatedPointSum)
		fmt.Println(user.RatedPointSumRank)
		fmt.Println(user.CreatedTime)
		fmt.Println(user.UpdatedTime)
	}
}

func getDate() string {
	const layout = "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(layout)
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "pass"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "user_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

// Users ユーザー情報のテーブル情報
type Users struct {
	Id                int
	Name              string `gorm:"column:name"`
	AcceptedCount     int    `gorm:"column:accepted_count"`
	AcceptedCountRank int    `gorm:"column:accepted_count_rank"`
	RatedPointSum     int    `gorm:"column:rated_point_sum"`
	RatedPointSumRank int    `gorm:"column:rated_point_sum_rank"`
	CreatedTime       string `gorm:"column:created_time" sql:"not null;type:date"`
	UpdatedTime       string `gorm:"column:updated_time" sql:"not null;type:date"`
}
