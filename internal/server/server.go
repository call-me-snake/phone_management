package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/call-me-snake/phone_management/internal/helpers"
	"github.com/call-me-snake/phone_management/internal/model"
	"github.com/call-me-snake/phone_management/internal/validate"
	"github.com/gorilla/mux"
	"github.com/hako/durafmt"
)

const BadRequestMessage = "Некорректные входные данные"
const InternalErrorMessage = "Внутренняя ошибка сервера"

const BanKeyPrefix = "Banned:"         //Префикс для сохранения ключа временного бана в формате Banned:UserId
const AttemptsLeftPrefix = "Attempts:" //Префикс для сохранения ключа количества оставшихся попыток в формате Attempts:UserId
const SuspendKeyPrefix = "Suspend:"    //Префикс для сохранения ключа временного таймаута в формате Suspend:UserId
const SmsKeyPrefix = "SmsKey:"         //Префикс для сохранения ключа временного кода в формате SmsKey:UserId:Phone

const defaultSmsKey = 5555 //Временный код для режима тестирования

//Connector - содержит роутер и адрес вызываемого сервиса
type Connector struct {
	router  *mux.Router
	address string
}

//New - Конструктор *Connector
func New(addr string) *Connector {
	c := &Connector{}
	c.router = mux.NewRouter()
	c.address = addr
	return c
}

func (c *Connector) executeHandlers(db model.IPhoneStorage, keyDb model.IKeyStorage) {
	c.router.HandleFunc("/ready", aliveHandler).Methods("GET")
	c.router.HandleFunc(`/getphone/{name:[\w-]+}`, getPhoneByName(db)).Methods("GET") //[\w-]+ regexp для символов [0-9A-Za-z_-]
	c.router.HandleFunc("/sendsms", sendSms(keyDb)).Methods("POST")
}

//Start запуск http сервера
func (c *Connector) Start(db model.IPhoneStorage, keyDb model.IKeyStorage) {
	c.executeHandlers(db, keyDb)
	http.ListenAndServe(c.address, c.router)
}

func aliveHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from phone management"))
}

//getPhoneByName - возвращает привязанный к пользователю номер телефона
func getPhoneByName(db model.IPhoneStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		name := mux.Vars(r)["name"]
		if name == "" {
			http.Error(w, BadRequestMessage, http.StatusBadRequest)
			return
		}

		result, err := db.GetPhone(name)
		if err != nil {
			http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
			log.Printf("getPhoneByName: %v", err)
			return
		}
		if result == nil {
			http.Error(w, fmt.Sprintf("Пользователя %s не найдено", name), http.StatusNotFound)
			return
		}
		if result.PhoneNumber[0] == ' ' {
			http.Error(w, fmt.Sprintf("Не задан номер пользователя %s", name), http.StatusNotFound)
			return
		}

		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
			log.Printf("getPhoneByName: %v", err)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(resp)
	}
}

//sendSms - отправляет пользователю смс с кодом
//пример нормального тела запроса: {"UserId":"user1","Phone":"71234567890"}
func sendSms(keyDb model.IKeyStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.SendSmsRequestJson{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil || user.UserId == "" {
			http.Error(w, BadRequestMessage, http.StatusBadRequest)
			return
		}
		if user.PhoneNumber, err = validate.ValidatePhone(user.PhoneNumber); err != nil {
			http.Error(w, "Некорректный формат телефона", http.StatusBadRequest)
			log.Printf("sendSms: %s", err.Error())
			return
		}

		//проверка превышения количества попыток в сутки
		userAttemptsLeft := AttemptsLeftPrefix + user.UserId
		attempts, err := keyDb.GetIntValueByKey(userAttemptsLeft)
		if err != nil {
			http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
			log.Printf("sendSms: %s", err.Error())
			return
		}

		if attempts != nil && *attempts <= 0 {
			http.Error(w, "Количество запросов смс в сутки было превышено.", http.StatusForbidden)
			return
		}

		//проверка бана пользователя. Пользователь банится от запросов при смене номера на час
		//TODO вписать переменную окр
		userBannedKey := BanKeyPrefix + user.UserId
		banned, err := keyDb.GetIntValueByKey(userBannedKey)
		if err != nil {
			http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
			log.Printf("sendSms: %s", err.Error())
			return
		}
		if banned != nil {
			timeout, err := returnTimeoutInString(userBannedKey, keyDb)
			if err != nil {
				http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
				log.Printf("sendSms: %s", err.Error())
				return
			}
			http.Error(w, fmt.Sprintf("Количество попыток ввода было превышено. Получение кода будет доступно через %s", timeout), http.StatusForbidden)
			return
		}

		//проверка ограничения пользователя. Ограничение накладывается при получении смс на минуту
		//TODO вписать переменную окр
		userSuspendedKey := SuspendKeyPrefix + user.UserId
		suspended, err := keyDb.GetIntValueByKey(userSuspendedKey)
		if err != nil {
			http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
			log.Printf("sendSms: %s", err.Error())
			return
		}
		if suspended != nil {
			timeout, err := returnTimeoutInString(userSuspendedKey, keyDb)
			if err != nil {
				http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
				log.Printf("sendSms: %s", err.Error())
				return
			}
			http.Error(w, fmt.Sprintf("Повторный запрос кода будет доступен через %s", timeout), http.StatusForbidden)
			return
		}

		//установка ограничения
		err = keyDb.SetTempIntKey(userSuspendedKey, 1, model.SuspendTimeout)
		if err != nil {
			http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
			log.Printf("sendSms: %s", err.Error())
			return
		}

		//установка ключа количества попыток, или его уменьшение
		var attemptsLeft int64
		switch {
		case attempts != nil:
			attemptsLeft, err = keyDb.DecrKey(userAttemptsLeft)
			if err != nil {
				http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
				log.Printf("sendSms: %s", err.Error())
				return
			}
			//если attemptsLeft<0, значит что-то пошло не так. Оно не должно быть <0 в нормальных условиях. Логирую
			if attemptsLeft < 0 {
				http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
				log.Printf("sendSms: attemptsLeft < 0, attemptsLeft = %d (Результат функции IKeyStorage.DecrKey не должен становиться < 0)", attemptsLeft)
				return
			}
		case attempts == nil:
			attemptsLeft = int64(model.TriesPerDay - 1)
			err = keyDb.SetTempIntKeyOnTimeStamp(userAttemptsLeft, model.TriesPerDay-1, helpers.GetNextDayDate())
			if err != nil {
				http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
				log.Printf("sendSms: %s", err.Error())
				return
			}
		}
		//имитация отправки сообщения
		err = smsSender(*user, keyDb)
		if err != nil {
			http.Error(w, InternalErrorMessage, http.StatusInternalServerError)
			log.Printf("sendSms: %s", err.Error())
			return
		}

		respMessage := model.SendSmsResponseJson{
			SendSmsRequestJson: *user,
			CodeLifeTime:       model.SmsKeyLifeSpan.String(),
			AttemptsLeft:       attemptsLeft,
		}

		resp, _ := json.Marshal(respMessage)
		w.Header().Set("content-type", "application/json")
		w.Write(resp)
	}
}

//returnTimeoutInString - возвращает таймаут в формате строки для дальнейшего вывода в тексте ошибки
func returnTimeoutInString(key string, db model.IKeyStorage) (string, error) {
	timeout, err := db.GetKeyLifeRest(key)
	if err != nil {
		return "", fmt.Errorf("server:returnTimeoutInString %s", err.Error())
	}
	if timeout == nil {
		return "", errors.New("server:returnTimeoutInString; Expected not nil timeout")
	}
	*timeout = timeout.Truncate(time.Second)
	timeoutstr := durafmt.Parse(*timeout).String()
	return timeoutstr, nil
}

//smsSender имитирует работу сервиса смс рассылки. Имеет 2 режима работы, зависящие от переменной model.InProduction
func smsSender(user model.SendSmsRequestJson, db model.IKeyStorage) error {
	userSmsKey := SmsKeyPrefix + user.UserId + ":" + user.PhoneNumber
	var smsKey int
	switch model.InProduction {
	case true:
		rand.Seed(time.Now().UnixNano())
		smsKey = helpers.RandomInRange(10000, 1000)
	case false:
		smsKey = defaultSmsKey
	}
	err := db.SetTempIntKey(userSmsKey, smsKey, model.SmsKeyLifeSpan)
	if err != nil {
		return fmt.Errorf("server:smsSender %s", err.Error())
	}
	log.Printf("server:smsSender: Пользователю %s отправлен код %d на номер %s", user.UserId, smsKey, user.PhoneNumber)
	return nil
}
