package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	t1value  = "+7-912-845-12-10"
	t1result = "79128451210"
	t2value  = "8 211 496 12-33"
	t2result = "72114961233"
	t3value  = "7 (123) 456 7 8 9-9"
	t3result = "71234567899"
	t4value  = "12345678901"
	t5value  = "82345678901212123"
	t6value  = "7 (123)a456b7c8 9-9"
	empty    = ""
)

func TestValidatePhoneT1(t *testing.T) {
	phone, err := ValidatePhone(t1value)
	assert.Equal(t, t1result, phone)
	assert.Nil(t, err)
}
func TestValidatePhoneT2(t *testing.T) {
	phone, err := ValidatePhone(t2value)
	assert.Equal(t, t2result, phone)
	assert.Nil(t, err)
}

func TestValidatePhoneT3(t *testing.T) {
	phone, err := ValidatePhone(t3value)
	assert.Equal(t, t3result, phone)
	assert.Nil(t, err)
}

func TestValidatePhoneT4(t *testing.T) {
	phone, err := ValidatePhone(t4value)
	assert.Equal(t, empty, phone)
	assert.Error(t, err)
}

func TestValidatePhoneT5(t *testing.T) {
	phone, err := ValidatePhone(t5value)
	assert.Equal(t, empty, phone)
	assert.Error(t, err)
}

func TestValidatePhoneT6(t *testing.T) {
	phone, err := ValidatePhone(t6value)
	assert.Equal(t, empty, phone)
	assert.Error(t, err)
}
