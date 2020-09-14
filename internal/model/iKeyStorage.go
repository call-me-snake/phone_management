package model

import "time"

//IKeyStorage - реализует методы хранилища временной информации по пользователям
type IKeyStorage interface {
	GetStringValueByKey(key string) (*string, error)
	GetIntValueByKey(key string) (*int, error)
	SetTempIntKey(key string, value int, timeout time.Duration) error
	SetTempIntKeyOnTimeStamp(key string, value int, timestamp time.Time) error
	GetKeyLifeRest(key string) (*time.Duration, error)
	DecrKey(key string) (int64, error)
	DelKey(key string) error
}
