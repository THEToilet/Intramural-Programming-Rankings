package main

type UserRank struct {
	Rate         int    `json:"Rate"`
	UserName     string `json:"UserName"`
	ContestCount int    `json:"ContestCount"`
}

func ranking() []*UserRank {

	users := loadFile("user.txt")

	var userRank *UserRank
	userRanks := make([]*UserRank, 0, 100)
	for _, user := range users {
		tmpUsersHistories := getAtCoderHistoryStruct(user)
		var userScore int
		var count int = 0
		for _, tmpUserHisory := range tmpUsersHistories {
			num, _ := tmpUserHisory.NewRating.Int64()
			userScore = int(num)
			count++
		}

		userRank = new(UserRank)
		userRank.Rate = userScore
		userRank.UserName = user
		userRank.ContestCount = count
		//	fmt.Println(userRank.Rate)
		//	fmt.Println(userRank.UserName)
		userRanks = append(userRanks, userRank)
	}

	return userRanks
}
