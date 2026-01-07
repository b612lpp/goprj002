package domain

import (
	"reflect"
	"testing"
)

func TestApply(t *testing.T) {

	tests := []struct {
		name string
		mr   MeterReading

		p, v     []int
		expected error
	}{
		{
			name:     "Добавляем показания газа в пустую базу",
			v:        []int{1},
			p:        []int{},
			mr:       MeterReading{counts: 1},
			expected: nil},
		{
			name:     "Добавляем показания света в пустую базу",
			v:        []int{1, 1},
			p:        []int{},
			mr:       MeterReading{counts: 2},
			expected: nil},
		{
			name:     "Добавляем отрицательные показания газа в пустую базу",
			v:        []int{-1},
			p:        []int{},
			mr:       MeterReading{counts: 1},
			expected: ErrValueLessThanZero},
		{
			name:     "Добавляем отрицательные показания света в пустую базу",
			v:        []int{-1, 1},
			p:        []int{},
			mr:       MeterReading{counts: 2},
			expected: ErrValueLessThanZero},
		{
			name:     "Добавляем корректные показания газа в не пустую базу",
			v:        []int{2},
			p:        []int{1},
			mr:       MeterReading{counts: 1},
			expected: nil},
		{
			name:     "Добавляем корректные показания света в не пустую базу",
			v:        []int{2, 2},
			p:        []int{1, 1},
			mr:       MeterReading{counts: 2},
			expected: nil},
		{
			name:     "Добавляем меньшие показания газа в не пустую базу",
			v:        []int{1},
			p:        []int{2},
			mr:       MeterReading{counts: 1},
			expected: ErrNewValueLessThanPrev},
		{
			name:     "Добавляем меньшие показания света в не пустую базу",
			v:        []int{1, 1},
			p:        []int{2, 2},
			mr:       MeterReading{counts: 2},
			expected: ErrNewValueLessThanPrev},
		{
			name:     "Добавляем больше показаний газа в не пустую базу",
			v:        []int{1, 2},
			p:        []int{1},
			mr:       MeterReading{counts: 1},
			expected: ErrValuesTypeMismatch},
		{
			name:     "Добавляем больше показаний света в не пустую базу",
			v:        []int{2, 2, 2},
			p:        []int{2, 2},
			mr:       MeterReading{counts: 2},
			expected: ErrValuesTypeMismatch},
		{
			name:     "Добавляем больше показаний газа в не пустую базу",
			v:        []int{},
			p:        []int{1},
			mr:       MeterReading{counts: 1},
			expected: ErrValuesTypeMismatch},
		{
			name:     "Добавляем меньше показаний света в не пустую базу",
			v:        []int{2},
			p:        []int{2, 2},
			mr:       MeterReading{counts: 2},
			expected: ErrValuesTypeMismatch},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.mr.Apply(tt.p, tt.v)

			if result == nil {
				if !reflect.DeepEqual(tt.v, tt.mr.values) {
					t.Errorf("%s подходящие значения не были применены", tt.name)
				}
			} else {
				if tt.mr.values != nil {
					t.Errorf("%s значения не должны были быть применены. записано %v", tt.name, tt.mr.values)
				}
			}

			if result != tt.expected {
				t.Errorf("%s failed got %v expexted %v", tt.name, result, tt.expected)
			}
		})
	}

}
