package main

import (
	"github.com/call-me-snake/phone_management/internal/postgrPhoneStorage"
	"github.com/call-me-snake/phone_management/internal/redisKeyStorage"
	"github.com/call-me-snake/phone_management/internal/server"
	"github.com/labstack/gommon/log"
)

var dbConn = "user=postgres password=example dbname=phone_storage sslmode=disable port=5432 host=localhost"

func main() {
	/*
		database, err := postgrPhoneStorage.New(dbConn)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = database.CreateOwner(model.PhoneOwner{Name: "test", PhoneNumber: ""})
		//err = database.DeleteOwner("tes1")
		//err = database.DeleteOwner("test")
		//owner, err := database.GetPhone("test")
		//fmt.Printf("%+v", owner)
		//if err != nil {
		//	log.Print(err.Error())
		//}
	*/
	keydb, err := redisKeyStorage.New("localhost:6379")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	database, err := postgrPhoneStorage.New(dbConn)
	if err != nil {
		log.Fatal(err.Error())
	}

	server := server.New(":8081")
	server.Start(database, keydb)
	/*
		err = db.SetTempIntKey("aaa", 123, 120*time.Second)
		if err != nil {
			log.Print(err)
			return
		}
		val, err := db.GetIntValueByKey("aaa")
		if val != nil {
			fmt.Println(*val)
		}
		if err != nil {
			log.Print(err)
			return
		}
		val, err = db.GetIntValueByKey("a")
		if err != nil {
			log.Print(err)
		}
		t := time.Now()
		t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		t2 := t1.AddDate(0, 0, 1)
		err = db.SetTempIntKeyOnTimeStamp("bbb", 12, t2)
		if err != nil {
			log.Print(err)
		}
	*/
}
