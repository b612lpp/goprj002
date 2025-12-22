package middleware

import "errors"

var (
	ErrBadCreds = errors.New("ошибка аутентификационных данных") //неполные или неверные данные аутентификации
	ErrDBConn   = errors.New("ошибка доступа к бд")
	ErrDBRead   = errors.New("ошибка чтения из бд")
)
