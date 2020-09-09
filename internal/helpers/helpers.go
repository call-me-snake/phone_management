package helpers

import (
	"math/rand"
	"time"
)

//GetNextDayDate - возвращает time.Time с 00.00.00 чч.мин.сс следующего дня
func GetNextDayDate() time.Time {
	now := time.Now()
	truncated := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	nextDay := truncated.AddDate(0, 0, 1)
	return nextDay
}

//RandomInRange - возвращает число в диапазоне от минимального до максимального
func RandomInRange(max, min int) int {
	return rand.Intn(max-min) + min
}
