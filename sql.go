package main

import (
 	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func test_sql() {
	// db接続
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	uusers := loadFile("user.txt")

	uuuu := []Users{}
	count := 1
	for _, user := range uusers {
		tmpInfo := getApi(user)
		// print(tmpInfo)
		//e := db.Where("Name = ?", user).Find(&uuuu)
		//if e != nil {
		db.First(&uuuu, count)
		count++
		// fmt.Println("#%v", tmpInfo)
		id, _ := tmpInfo.AcceptedCount.Int64()
		id1, _ := tmpInfo.AcceptedCountRank.Int64()
		id2, _ := tmpInfo.RatedPointSum.Int64()
		id3, _ := tmpInfo.RatedPointSumRank.Int64()
		db.Create(&Users{
			Name:              tmpInfo.UserId,
			AcceptedCount:     int(id),
			AcceptedCountRank: int(id1),
			RatedPointSum:     int(id2),
			RatedPointSumRank: int(id3),
			CreatedTime:       getDate(),
			UpdatedTime:       getDate(),
		})
		//}
		/*
			uuuu.Name = tmpInfo.UserId
			uuuu.AcceptedCount = int(id)
			uuuu.AcceptedCountRank = int(id1)
			uuuu.RatedPointSum = int(id2)
			uuuu.RatedPointSumRank = int(id3)
			uuuu.CreatedTime = getDate()
			uuuu.UpdatedTime = getDate()
			db.Save(&uuuu)*/
	}

	/*
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
	*/
}

func getUserInfo() string {

	// db接続
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	//データを格納する変数を定義
	users := []Users{}

	//全取得
	db.Find(&users)
	//	print(users)

	var result string
	//表示
	for _, user := range users {

	  fmt.Println(user)
		fmt.Println(user.Id)
		fmt.Println(user.Name)
		fmt.Println(user.AcceptedCount)
		fmt.Println(user.AcceptedCountRank)
		fmt.Println(user.RatedPointSum)
		fmt.Println(user.RatedPointSumRank)
		fmt.Println(user.CreatedTime)
		fmt.Println(user.UpdatedTime)

//	for _, user := range users {
		result += (fmt.Sprintf("%d",user.Id) +" "+ fmt.Sprintf("%s",user.Name) + " "+ fmt.Sprintf("%d",user.AcceptedCount) +" "+ fmt.Sprintf("%d",user.AcceptedCountRank) +" "+ fmt.Sprintf("%d",user.RatedPointSum) +" "+ fmt.Sprintf("%d",user.RatedPointSumRank) +" " + fmt.Sprintf("%s",user.CreatedTime) + "\n")
	}

//result += "unnko"

	return result

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

// Users ユーザー情報のテーブル情報(SQLの構造体)
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
