package validate

import (
	"errors"
	"regexp"
	"unicode"
)

const numberLen = 11

func ValidatePhone(input string) (string, error) {
	reg := regexp.MustCompile("[-() ]+")
	s := reg.ReplaceAllString(input, "")
	if s[0] == '+' {
		s = s[1:]
	}
	if (len(s) != numberLen) || !(s[0] == '7' || s[0] == '8') {
		return "", errors.New("ValidatePhone: Неверный формат строки")
	}
	for i := 1; i < len(s); i++ {
		if !unicode.IsNumber(rune(s[i])) {
			return "", errors.New("ValidatePhone: Неверный формат строки")
		}
	}
	if s[0] == '8' {
		s = "7" + s[1:]
	}
	return s, nil
}
