package fabric

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/b612lpp/goprj002/domain"
)

func TestEventFabric(t *testing.T) {

	gmr := domain.NewGasReading("me")
	gmr.Apply([]int{}, []int{1})
	gt := gmr.GetCreatedAt()
	emr := domain.NewEnReading("me")
	err := emr.Apply([]int{}, []int{1, 2})
	et := emr.GetCreatedAt()
	if err != nil {
		t.Fatal("не получили данные из домена")
	}

	tests := []struct {
		name     string
		ef       EventFormer
		mr       domain.MeterReading
		expected domain.EventuallyAppliedData
	}{
		{

			name:     "Счетчик газа",
			ef:       NewEventFabric(),
			mr:       gmr,
			expected: domain.EventuallyAppliedData{Owner: "me", MeterType: "_Gas", Values: []int{1}, CreatedAt: gt}},
		{

			name:     "Счетчик света",
			ef:       NewEventFabric(),
			mr:       emr,
			expected: domain.EventuallyAppliedData{Owner: "me", MeterType: "_Electro", Values: []int{1, 2}, CreatedAt: et}},
	}
	for _, tt := range tests {
		result := tt.ef.MakeEvent(tt.mr)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("%s failed. wanted %v got %v", tt.name, tt.expected, result)
		}
		fmt.Println(result)
	}
}
