package webservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type GameRecord struct {
	UserName string `json:"username"`
	Score    uint   `json:"score`
}

var GameDB *gorm.DB

func LoadGameDB() {
	var err error
	GameDB, err = gorm.Open(sqlite.Open("MiniGame.db"), &gorm.Config{})
	fmt.Printf("LoadGameDB")
	if err != nil {
		log.Fatalln(err)
	}

	res := GameDB.Find(&GameRecordAry)
	GameRecordAryLen = int32(res.RowsAffected)
	fmt.Printf("Done\n")
}

func AddGameRecord(record GameRecord) {
	if GameDB != nil {
		res := GameDB.Create(&record)

		if res.Error != nil {
			fmt.Println(res.Error)
		}
	}
}

func GetTop10GameRecord() (records []GameRecord, success bool) {
	success = false

	if GameDB != nil {
		if err := GameDB.Model(&GameRecord{}).Order("score desc").Limit(10).Find(&records).Error; err == nil {
			success = true
		}
	}

	return
}

func HttpUploadRecord(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var record GameRecord
		err := json.NewDecoder(req.Body).Decode(&record)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	default:
	}
}
