package application

import "errors"

var (
	ErrGetDataFromDb   = errors.New("ошибка получения данных из бд")
	ErrValueValidation = errors.New("значение не может быть меньше предыдущего")
	ErrUnknown         = errors.New("неизвестная ошибка")
)
