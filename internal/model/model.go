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

type envs struct {
	ServerAddress       string `long:"http" env:"SERVER" description:"address of microservice" default:":8081"`
	PhoneStorageConn    string `long:"pstconn" env:"PHONE_STORAGE" description:"Connection string to phone storage database" default:"user=postgres password=example dbname=phone_storage sslmode=disable port=5432 host=localhost"`
	KeyStorageConn      string `long:"kstconn" env:"KEY_STORAGE" description:"Connection string to key storage database" default:":6379"`
	SuspendTimeout      string `long:"stimeout" env:"SUSPEND_TIMEOUT" description:"Suspend timeout after getting key" default:"1m"`
	SmsKeyLifeSpan      string `long:"smskeylife" env:"SMSKEY_LIFESPAN" description:"Sms key lifespan after getting key" default:"15m"`
	AttemptsKeyLifeSpan string `long:"attemptskeylife" env:"ATTEMPTS_KEY_LIFESPAN" description:"Attempts key lifespan after failing sms" default:"15m"`
	BanKeyLifeSpan      string `long:"bankeylife" env:"BAN_KEY_LIFESPAN" description:"Ban key lifespan after successfully changing phone" default:"1h"`
	TriesPerDay         string `long:"triesperday" env:"TRIES_PER_DAY" description:"Tries of getting sms per day" default:"3"`
	AttemptsOfInput     string `long:"iattempts" env:"INPUT_ATTEMPTS" description:"Attempts of input sms code per AttemptsKeyLifeSpan" default:"3"`
	InProduction        bool   `long:"inprod" env:"IN_PROD" description:"Shows whether service is in test mode or not" default:"false"`
}
