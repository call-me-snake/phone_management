package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jessevdk/go-flags"
)

var (
	C = &Config{}
)

//envs - приватная структура, получает переменные окружения
type envs struct {
	ServerAddress       string `long:"http" env:"SERVER" description:"address of microservice" default:":8081"`
	PhoneStorageConn    string `long:"pstconn" env:"PHONE_STORAGE" description:"Connection string to phone storage database" default:"user=postgres password=example dbname=phone_storage sslmode=disable port=5432 host=localhost"`
	KeyStorageConn      string `long:"kstconn" env:"KEY_STORAGE" description:"Connection string to key storage database" default:":6379"`
	SuspendTimeout      string `long:"stimeout" env:"SUSPEND_TIMEOUT" description:"Suspend timeout after getting key" default:"1m"`
	SmsKeyLifeSpan      string `long:"smskeylife" env:"SMSKEY_LIFESPAN" description:"Sms key lifespan after getting key" default:"15m"`
	AttemptsKeyLifeSpan string `long:"attemptskeylife" env:"ATTEMPTS_KEY_LIFESPAN" description:"Attempts key lifespan after failing sms" default:"1h"`
	BanKeyLifeSpan      string `long:"bankeylife" env:"BAN_KEY_LIFESPAN" description:"Ban key lifespan after successfully changing phone" default:"1h"`
	TriesPerDay         string `long:"triesperday" env:"TRIES_PER_DAY" description:"Tries of getting sms per day" default:"3"`
	AttemptsOfInput     string `long:"iattempts" env:"INPUT_ATTEMPTS" description:"Attempts of input sms code per AttemptsKeyLifeSpan" default:"3"`
	InProduction        bool   `long:"inprod" env:"IN_PROD" description:"Shows whether service is in test mode or not"`
}

//Config - публичная структура, хранит проверенные переменные окружения
type Config struct {
	ServerAddress       string        //Адрес сервера
	PhoneStorageConn    string        //Строка соединения к хранилищу пользователей
	KeyStorageConn      string        //Строка соединения к хранилищу ключей
	SuspendTimeout      time.Duration //SuspendTimeout - таймаут после отправления смс в server.sendSms
	SmsKeyLifeSpan      time.Duration //SmsKeyLifeSpan - время жизни смс ключа
	AttemptsKeyLifeSpan time.Duration //AttemptsKeyLifeSpan - время жизни ключа количества попыток
	BanKeyLifeSpan      time.Duration //BanKeyLifeSpan - время жизни ключа бана после смены номера
	TriesPerDay         int           //TriesPerDay - максимальное количество получаемых смс в сутки
	AttemptsOfInput     int           //AttemptsOfInput - максимальное количество попыток ввода кода в промежуток времени AttemptsKeyLifeSpan
	InProduction        bool          //InProduction - булевая переменная, влияет на работу функции server.SmsSender. В случае false SmsSender все время посылает одинаковый код
}

//GetConfig - получает переменные окружения с помощью приватной структуры envs и проверяет их
func GetConfig() error {
	e := envs{}
	var err error
	parser := flags.NewParser(&e, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return fmt.Errorf("GetConfig: %v", err)
	}
	C.ServerAddress = e.ServerAddress
	C.PhoneStorageConn = e.PhoneStorageConn
	C.KeyStorageConn = e.KeyStorageConn
	C.SuspendTimeout, err = time.ParseDuration(e.SuspendTimeout)
	if err != nil {
		return fmt.Errorf("GetConfig: %v", err)
	}
	C.SmsKeyLifeSpan, err = time.ParseDuration(e.SmsKeyLifeSpan)
	if err != nil {
		return fmt.Errorf("GetConfig: %v", err)
	}
	C.AttemptsKeyLifeSpan, err = time.ParseDuration(e.AttemptsKeyLifeSpan)
	if err != nil {
		return fmt.Errorf("GetConfig: %v", err)
	}
	C.BanKeyLifeSpan, err = time.ParseDuration(e.BanKeyLifeSpan)
	if err != nil {
		return fmt.Errorf("GetConfig: %v", err)
	}
	C.TriesPerDay, err = strconv.Atoi(e.TriesPerDay)
	if err != nil {
		return fmt.Errorf("GetConfig: %v", err)
	}
	C.AttemptsOfInput, err = strconv.Atoi(e.AttemptsOfInput)
	if err != nil {
		return fmt.Errorf("GetConfig: %v", err)
	}
	C.InProduction = e.InProduction
	return nil
}

//GetConfigForTest - используется для установки конфига в тестах
func GetConfigForTest() {
	C.SuspendTimeout = time.Minute
	C.SmsKeyLifeSpan = 15 * time.Minute
	C.AttemptsKeyLifeSpan = time.Hour
	C.BanKeyLifeSpan = time.Hour
	C.TriesPerDay = 3
	C.AttemptsOfInput = 3
	C.InProduction = false
}

func SetNilConfig() {
	C = &Config{}
}
