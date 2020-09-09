package model

import "time"

var (
	//SuspendTimeout - таймаут после отправления смс в server.sendSms
	SuspendTimeout time.Duration = time.Minute * 1

	//SmsKeyLifeSpan - время жизни смс ключа
	SmsKeyLifeSpan time.Duration = time.Minute * 15

	//TriesPerDay - максимальное количество получаемых смс в сутки
	TriesPerDay = 3

	//InProduction - булевая переменная, влияет на работу функции server.SmsSender. В случае false SmsSender все время посылает одинаковый код
	InProduction = false
)
