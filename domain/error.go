package domain

import "errors"

var (
	ErrValueToAdd           = errors.New("некорректное значение показаний для добавления")
	ErrValuesMismatch       = errors.New("переданные данные не соответствуют формату")
	ErrNewValueLessThanPrev = errors.New("новое значение меньше чем предыдущее")
	ErrValueLessThanZero    = errors.New("новое значение не может быть меньше нуля")
	ErrEmptyValues          = errors.New("новые данные не могут быть пустыми")
)
