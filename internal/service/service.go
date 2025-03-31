package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

const (
	sign1 = "."
	sign2 = "-"
)

// Функция конвертирует код Морзе в текст
func ToConvert(text string) string {

	if IsTextMorse(text) {
		//2 - и наоборот — если был передан код Морзе, функция должна переконвертировать его в обычный текст и вернуть
		return morse.ToText(text)
	}
	//1 -Если передан обычный текст, функция должна переконвертировать его в код Морзе и вернуть
	return morse.ToMorse(text)
}

// Функция проверяет является ли строка текстом или кодом Морзе
func IsTextMorse(text string) bool {

	if strings.Contains(text, sign1) || strings.Contains(text, sign2) {
		return true
	}
	return false
}
