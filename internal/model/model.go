package model

//PhoneOwner - содержит информацию: имя владельца телефона и номер
type PhoneOwner struct {
	Name        string `gorm:"primaryKey;column:name" json:"UserId"`
	PhoneNumber string `gorm:"column:phone_number;size:11" json:"Phone"`
}

//SendSmsRequestJson - структура для записи получаемых данных в ручке server.sendSms
type SendSmsRequestJson struct {
	UserId      string `json:"UserId"`
	PhoneNumber string `json:"Phone"`
}

//SendSmsResponseJson - структура для записи возвращаемых данных в ручке server.sendSms
type SendSmsResponseJson struct {
	SendSmsRequestJson
	CodeLifeTime string `json:"CodeLifeTime"`
	RequestsLeft int64  `json:"RequestsLeft"`
}

//AttachNewPhoneRequestJson - структура для записи получаемых данных в ручке server.attachNewPhone
type AttachNewPhoneRequestJson struct {
	SendSmsRequestJson
	Code int `json:"Code"`
}

//AttachNewPhoneResponseJson - структура для записи возвращаемых данных в ручке server.attachNewPhone
type AttachNewPhoneResponseJson struct {
	SendSmsRequestJson
}
