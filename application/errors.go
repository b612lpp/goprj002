package application

import "errors"

var (
	ErrGetDataFromDb   = errors.New("ошибка получения данных из бд")
	ErrValueValidation = errors.New("значение не может быть меньше предыдущего")
	ErrLowerZero       = errors.New("значение не может быть меньше 0")
	ErrUnknown         = errors.New("неизвестная ошибка")
)
