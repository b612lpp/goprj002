package domain

import (
	"testing"
)

func TestIsValidComparedTo(t *testing.T) {

	tests := []struct {
		name     string
		mr       MeterReading
		p        []int
		expected error
	}{
		{
			name:     "Новые показания больше чем предыдущие",
			mr:       MeterReading{Values: []int{1, 1}},
			p:        []int{0, 0},
			expected: nil},
		{
			name:     "Новые показания такие же",
			mr:       MeterReading{Values: []int{1, 1}},
			p:        []int{1, 1},
			expected: nil,
		},
		{
			name:     "Оба показания меньше",
			mr:       MeterReading{Values: []int{1, 1}},
			p:        []int{2, 2},
			expected: ErrNewValueLessThanPrev,
		},
		{
			name:     "Первое показание меньше",
			mr:       MeterReading{Values: []int{1, 1}},
			p:        []int{2, 1},
			expected: ErrNewValueLessThanPrev,
		},
		{
			name:     "Второе показание меньше",
			mr:       MeterReading{Values: []int{1, 1}},
			p:        []int{1, 2},
			expected: ErrNewValueLessThanPrev,
		},
		{
			name:     "Новые показания меньше 0",
			mr:       MeterReading{Values: []int{-1, -1}},
			p:        []int{2, 2},
			expected: ErrValueLessThanZero,
		},
		{
			name:     "Массивы разной длинны. Новый больше",
			mr:       MeterReading{Values: []int{1, 1}},
			p:        []int{2},
			expected: ErrValuesMismatch,
		},
		{
			name:     "Массивы разной длинны. Старый больше",
			mr:       MeterReading{Values: []int{3}},
			p:        []int{2, 2},
			expected: ErrValuesMismatch,
		}, {
			name:     "Входящий массив пустой",
			mr:       MeterReading{Values: []int{}},
			p:        []int{2, 2},
			expected: ErrEmptyValues,
		},
		{
			name:     "Массив предыдущих значений пустой",
			mr:       MeterReading{Values: []int{2, 2}},
			p:        []int{},
			expected: nil,
		},
		{
			name:     "Массив предыдущих значений отрицательный",
			mr:       MeterReading{Values: []int{2, 2}},
			p:        []int{-1, -1},
			expected: nil,
		},
		{
			name:     "Массивы оба отрицательные",
			mr:       MeterReading{Values: []int{-3, -3}},
			p:        []int{-4, -42},
			expected: ErrValueLessThanZero,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.mr.IsValid(tt.p)

			if result != tt.expected {
				t.Errorf("%s failed got %v expexted %v", tt.name, result, tt.expected)
			}
		})
	}

}
