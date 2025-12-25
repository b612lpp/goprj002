package repository

import "errors"

var (
	ErrDBConnection = errors.New("ошибка получения данных") //неполные или неверные данные аутентификации
	ErrEmptyData    = errors.New("нет данных")
)
