package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/call-me-snake/phone_management/internal/model"
	mock_model "github.com/call-me-snake/phone_management/internal/model/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var testOwnerName = "tester"
var testPhoneNumber = "79991234567"

var testPhoneOwner = model.PhoneOwner{
	Name:        testOwnerName,
	PhoneNumber: testPhoneNumber,
}

var testPhoneOwner1 = model.PhoneOwner{
	Name:        testOwnerName,
	PhoneNumber: "           ",
}

//TestAliveHandler ...
func TestAliveHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ready", nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(aliveHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, []byte("Hello from phone management"), rr.Body.Bytes())
}

//TestGetPhoneByName - тест выполнения
func TestGetPhoneByName(t *testing.T) {
	//определил mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockdb := mock_model.NewMockIPhoneStorage(ctrl)
	mockdb.
		EXPECT().
		GetPhone(testOwnerName).
		Return(&testPhoneOwner, nil)

	//делаю с помощью mux.NewRouter() из-за mux.Vars
	router := mux.NewRouter()
	router.HandleFunc(`/getphone/{name:[\w-]+}`, getPhoneByName(mockdb)).Methods("GET")
	req, err := http.NewRequest("GET", fmt.Sprintf("/getphone/%s", testOwnerName), nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	res, _ := json.Marshal(testPhoneOwner)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, res, rr.Body.Bytes())
}

//TestGetPhoneByNameFailInternal - тест ошибки доступа к бд
func TestGetPhoneByNameFailInternal(t *testing.T) {
	//определил mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockdb := mock_model.NewMockIPhoneStorage(ctrl)
	mockdb.
		EXPECT().
		GetPhone(testOwnerName).
		Return(nil, errors.New("error"))

	//делаю с помощью mux.NewRouter() из-за mux.Vars
	router := mux.NewRouter()
	router.HandleFunc(`/getphone/{name:[\w-]+}`, getPhoneByName(mockdb)).Methods("GET")
	req, err := http.NewRequest("GET", fmt.Sprintf("/getphone/%s", testOwnerName), nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

//TestGetPhoneByNameNilResult тест отсутствия пользователя
func TestGetPhoneByNameNilResult(t *testing.T) {
	//определил mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockdb := mock_model.NewMockIPhoneStorage(ctrl)
	mockdb.
		EXPECT().
		GetPhone(testOwnerName).
		Return(nil, nil)

	//делаю с помощью mux.NewRouter() из-за mux.Vars
	router := mux.NewRouter()
	router.HandleFunc(`/getphone/{name:[\w-]+}`, getPhoneByName(mockdb)).Methods("GET")
	req, err := http.NewRequest("GET", fmt.Sprintf("/getphone/%s", testOwnerName), nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

//TestGetPhoneByNameEmptyPhone тест отсутствия телефона у пользователя
func TestGetPhoneByNameEmptyPhone(t *testing.T) {
	//определил mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockdb := mock_model.NewMockIPhoneStorage(ctrl)
	mockdb.
		EXPECT().
		GetPhone(testOwnerName).
		Return(&testPhoneOwner1, nil)

	//делаю с помощью mux.NewRouter() из-за mux.Vars
	router := mux.NewRouter()
	router.HandleFunc(`/getphone/{name:[\w-]+}`, getPhoneByName(mockdb)).Methods("GET")
	req, err := http.NewRequest("GET", fmt.Sprintf("/getphone/%s", testOwnerName), nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

//TestSendSms - тест успешной отправки смс
func TestSendSms(t *testing.T) {
	model.GetConfigForTest()
	defer model.SetNilConfig()

	requestsLeftKey := "Requests:Test"
	requestsLeft := 2
	var requestsLeftDecr int64 = 1
	userBannedKey := "Banned:Test"
	userSuspendedKey := "Suspend:Test"
	userSmsKey := "SmsKey:Test:71234567890"
	requestModel := model.SendSmsRequestJson{UserId: "Test", PhoneNumber: "71234567890"}
	requestBody, _ := json.Marshal(requestModel)
	res, _ := json.Marshal(model.SendSmsResponseJson{SendSmsRequestJson: requestModel, CodeLifeTime: model.C.SmsKeyLifeSpan.String(), RequestsLeft: requestsLeftDecr})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockkeydb := mock_model.NewMockIKeyStorage(ctrl)
	mockkeydb.
		EXPECT().
		GetIntValueByKey(requestsLeftKey).
		Return(&requestsLeft, nil)

	mockkeydb.
		EXPECT().
		GetIntValueByKey(userBannedKey).
		Return(nil, nil)

	mockkeydb.
		EXPECT().
		GetIntValueByKey(userSuspendedKey).
		Return(nil, nil)

	mockkeydb.
		EXPECT().
		DecrKey(requestsLeftKey).
		Return(requestsLeftDecr, nil)

	mockkeydb.
		EXPECT().
		SetTempIntKey(userSuspendedKey, 1, model.C.SuspendTimeout).
		Return(nil)

	mockkeydb.
		EXPECT().
		SetTempIntKey(userSmsKey, 5555, model.C.SmsKeyLifeSpan).
		Return(nil)

	req, err := http.NewRequest("POST", "/sendsms", bytes.NewReader(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sendSms(mockkeydb))
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, res, rr.Body.Bytes())
}

//TestAttachNewPhone - тест успешной привязки телефона

func TestAttachNewPhone(t *testing.T) {
	model.GetConfigForTest()
	defer model.SetNilConfig()

	userAttemptsLeftKey := "Attempts:Test"
	attemptsLeft := 2
	userSmsKey := "SmsKey:Test:71234567890"
	key := 5555
	updatePhone := model.PhoneOwner{Name: "Test", PhoneNumber: "71234567890"}
	userBanKey := "Banned:Test"
	userModel := model.SendSmsRequestJson{UserId: "Test", PhoneNumber: "71234567890"}
	requestBody, _ := json.Marshal(model.AttachNewPhoneRequestJson{SendSmsRequestJson: userModel, Code: 5555})
	res, _ := json.Marshal(model.AttachNewPhoneResponseJson{SendSmsRequestJson: userModel})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockkeydb := mock_model.NewMockIKeyStorage(ctrl)
	mockuserdb := mock_model.NewMockIPhoneStorage(ctrl)

	mockkeydb.
		EXPECT().
		GetIntValueByKey(userAttemptsLeftKey).
		Return(&attemptsLeft, nil)

	mockkeydb.
		EXPECT().
		GetIntValueByKey(userSmsKey).
		Return(&key, nil)

	mockuserdb.
		EXPECT().
		UpdatePhone(updatePhone).
		Return(true, nil)

	mockkeydb.
		EXPECT().
		SetTempIntKey(userBanKey, 1, model.C.BanKeyLifeSpan).
		Return(nil)

	mockkeydb.EXPECT().DelKey(userAttemptsLeftKey).Return(nil)
	mockkeydb.EXPECT().DelKey(userSmsKey).Return(nil)

	req, err := http.NewRequest("POST", "/attachphone", bytes.NewReader(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(attachNewPhone(mockuserdb, mockkeydb))
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, res, rr.Body.Bytes())

}
