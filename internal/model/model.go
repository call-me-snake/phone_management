package model

//PhoneOwner - содержит информацию: имя владельца телефона и номер
type PhoneOwner struct {
	Name        string `gorm:"primaryKey;column:name" json:"UserId"`
	PhoneNumber string `gorm:"column:phone_number;size:11" json:"Phone"`
}

type SendSmsRequestJson struct {
	UserId      string `json:"UserId"`
	PhoneNumber string `json:"Phone"`
}

type SendSmsResponseJson struct {
	SendSmsRequestJson
	CodeLifeTime string `json:"CodeLifeTime"`
	AttemptsLeft int64  `json:"AttemptsLeft"`
}
