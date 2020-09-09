package server

import (
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

//TestGetPhoneByName - тест ошибки доступа к бд
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
