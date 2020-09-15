package main

import (
	"github.com/call-me-snake/phone_management/internal/model"
	"github.com/call-me-snake/phone_management/internal/postgrPhoneStorage"
	"github.com/call-me-snake/phone_management/internal/redisKeyStorage"
	"github.com/call-me-snake/phone_management/internal/server"
	"github.com/labstack/gommon/log"
)

func main() {
	log.Print("Started")
	//Устанавливаем значения переменных окружения
	err := model.GetConfig()
	if err != nil {
		log.Print(err.Error())
		return
	}
	//подключаем хранилище ключей
	keydb, err := redisKeyStorage.New(model.C.KeyStorageConn)
	if err != nil {
		log.Print(err.Error())
		return
	}
	//подключаем хранилище пользователей
	database, err := postgrPhoneStorage.New(model.C.PhoneStorageConn)
	if err != nil {
		log.Print(err.Error())
		return
	}
	//разворачиваем сервер
	server := server.New(model.C.ServerAddress)
	err = server.Start(database, keydb)
	if err != nil {
		log.Print(err.Error())
		return
	}
}
