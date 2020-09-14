package model

import "time"

var (
	//SuspendTimeout - таймаут после отправления смс в server.sendSms
	SuspendTimeout time.Duration = time.Minute * 1

	//SmsKeyLifeSpan - время жизни смс ключа
	SmsKeyLifeSpan time.Duration = time.Minute * 15

	//AttemptsKeyLifeSpan - время жизни ключа количества попыток
	AttemptsKeyLifeSpan time.Duration = time.Hour * 1

	//BanKeyLifeSpan - время жизни ключа бана после смены номера
	BanKeyLifeSpan time.Duration = time.Hour * 1

	//TriesPerDay - максимальное количество получаемых смс в сутки
	TriesPerDay = 3

	//AttemptsOfInput - максимальное количество попыток ввода кода в промежуток времени AttemptsKeyLifeSpan
	AttemptsOfInput = 3

	//InProduction - булевая переменная, влияет на работу функции server.SmsSender. В случае false SmsSender все время посылает одинаковый код
	InProduction = false
)
