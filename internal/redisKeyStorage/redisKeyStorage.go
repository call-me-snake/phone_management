package redisKeyStorage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/call-me-snake/phone_management/internal/model"
	"github.com/go-redis/redis"
)

//sleepDurationInSec - время пинга функции checkConnection в секундах
const sleepDurationInSec = 5

var ctx = context.Background()

//storage ...
type storage struct {
	rdb     *redis.Client
	address string
}

//New возвращает объект интерфейса IKeyStorage (storage)
func New(address string) (model.IKeyStorage, error) {
	db := &storage{}
	db.rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	err := db.ping()
	if err != nil {
		return nil, fmt.Errorf("redisKeyStorage.New: %s", err.Error())
	}
	return db, nil
}

//ping - внутренняя функция проверки соединения
func (db *storage) ping() error {
	pong, err := db.rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("redisKeyStorage.ping: %v", err)
	}
	if pong != "PONG" {
		return fmt.Errorf("redisKeyStorage.ping: pong = %s", pong)
	}
	return nil
}

//checkConnection - внутренняя функция постоянной проверки соединения
func (db *storage) checkConnection() {
	go func() {
		for {
			err := db.ping()
			if err != nil {
				log.Printf("redisKeyStorage.checkConnection: no connection: %s", err.Error())

				tempDb := redis.NewClient(&redis.Options{
					Addr:     db.address,
					Password: "",
					DB:       0,
				})
				if pong, err := tempDb.Ping(context.Background()).Result(); err != nil || pong != "PONG" {
					log.Printf("redisKeyStorage.checkConnection: coud not establish connection: err=%s,ping=%s", err.Error(), pong)
				} else {
					db.rdb = tempDb
				}
			}
			time.Sleep(time.Second * sleepDurationInSec)
		}
	}()
}

//GetStringValueByKey - получить строковое значение по ключу
func (db *storage) GetStringValueByKey(key string) (*string, error) {
	get := db.rdb.Get(ctx, key)
	if err := get.Err(); err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("redisKeyStorage.GetStringValueByKey: %v", err)
	}
	s := get.Val()
	return &s, nil
}

//GetIntValueByKey - получить значение типа int по ключу
func (db *storage) GetIntValueByKey(key string) (*int, error) {
	i, err := db.rdb.Get(ctx, key).Int()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("redisKeyStorage.GetIntValueByKey: %v", err)
	}
	return &i, nil
}

//SetTempIntKey - установить временный ключ типа int, уничтожающийся по истечению timeout
func (db *storage) SetTempIntKey(key string, value int, timeout time.Duration) error {
	err := db.rdb.Set(ctx, key, value, timeout).Err()
	if err != nil {
		return fmt.Errorf("redisKeyStorage.SetTempIntKey: %v", err)
	}
	return nil
}

//SetTempIntKeyOnTimeStamp - установить временный ключ типа int, уничтожающийся с наступлением timestamp
func (db *storage) SetTempIntKeyOnTimeStamp(key string, value int, timestamp time.Time) error {
	err := db.rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return fmt.Errorf("redisKeyStorage.SetTempIntKeyOnTimeStamp: %v", err)
	}
	db.rdb.ExpireAt(ctx, key, timestamp).Err()
	if err != nil {
		return fmt.Errorf("redisKeyStorage.SetTempIntKeyOnTimeStamp: %v", err)
	}
	return nil
}

//GetKeyLifeRest - получить остаток жизни ключа
func (db *storage) GetKeyLifeRest(key string) (*time.Duration, error) {
	timeRest, err := db.rdb.TTL(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("redisKeyStorage.GetKeyLifeRest: %v", err)
	}
	if timeRest < 0 {
		return nil, nil
	}
	return &timeRest, nil
}

//DecrKey - уменьшает значение ключа на 1
func (db *storage) DecrKey(key string) (int64, error) {
	i, err := db.rdb.Decr(ctx, key).Result()
	if err != nil {
		return i, fmt.Errorf("redisKeyStorage.DecrKey: %v", err)
	}
	return i, nil
}
