package postgrPhoneStorage

import (
	"fmt"
	"log"
	"time"

	"github.com/call-me-snake/phone_management/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//sleepDurationInSec - время пинга функции checkConnection в секундах
const sleepDurationInSec = 5

//storage ...
type storage struct {
	database *gorm.DB
	address  string
}

//New возвращает объект интерфейса IPhoneStorage (storage)
func New(adress string) (model.IPhoneStorage, error) {
	var err error
	db := &storage{}
	db.address = adress
	db.database, err = gorm.Open("postgres", adress)
	if err != nil {
		return nil, fmt.Errorf("postgrPhoneStorage.New: %v", err)
	}

	err = db.ping()
	if err != nil {
		return nil, fmt.Errorf("postgrPhoneStorage.New: %s", err.Error())
	}
	db.checkConnection()

	return db, nil
}

//ping (internal)
func (db *storage) ping() error {
	//db.database.LogMode(true)
	result := struct {
		Result int
	}{}

	err := db.database.Raw("select 1+1 as result").Scan(&result).Error
	if err != nil {
		return fmt.Errorf("postgrPhoneStorage.ping: %v", err)
	}
	if result.Result != 2 {
		return fmt.Errorf("postgrPhoneStorage.ping: incorrect result!=2 (%d)", result.Result)
	}
	return nil
}

//checkConnection (internal)
func (db *storage) checkConnection() {
	go func() {
		for {
			err := db.ping()
			if err != nil {

				log.Printf("postgrPhoneStorage.checkConnection: no connection: %s", err.Error())
				tempDb, err := gorm.Open("postgres", db.address)

				if err != nil {
					log.Printf("postgrPhoneStorage.checkConnection: could not establish connection: %v", err)
				} else {
					db.database = tempDb
				}
			}
			time.Sleep(sleepDurationInSec * time.Second)
		}
	}()
}

//GetPhone реализует функцию интерфейса IPhoneStorage
func (db *storage) GetPhone(owner string) (*model.PhoneOwner, error) {
	result := &model.PhoneOwner{}
	query := db.database.Where("name = ?", owner).First(result)
	if query.Error != nil {
		return nil, fmt.Errorf("GetPhone: %v", query.Error)
	}
	if query.RowsAffected == 0 {
		return nil, nil
	}
	return result, nil
}

//CreateOwner реализует функцию интерфейса IPhoneStorage
func (db *storage) CreateOwner(data model.PhoneOwner) error {
	query := db.database.Create(&data)
	if query.Error != nil {
		return fmt.Errorf("CreateOwner: %v", query.Error)
	}
	return nil
}

//UpdatePhone реализует функцию интерфейса IPhoneStorage
func (db *storage) UpdatePhone(data model.PhoneOwner) (isUpdated bool, err error) {
	query := db.database.Model(&model.PhoneOwner{}).Where("name = ?", data.Name).Update(data)
	if query.Error != nil {
		return false, fmt.Errorf("UpdatePhone: %v", query.Error)
	}
	if query.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

//DeleteOwner реализует функцию интерфейса IPhoneStorage
func (db *storage) DeleteOwner(owner string) error {
	query := db.database.Where("name = ?", owner).Delete(&model.PhoneOwner{})
	if query.Error != nil {
		return fmt.Errorf("DeleteOwner: %v", query.Error)
	}
	return nil
}
