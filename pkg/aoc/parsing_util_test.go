package aoc

import (
	"reflect"
	"testing"
)

func Test_IntsFromString(t *testing.T) {
	input := "1,2,-3,4,5"
	expected := []int64{1, 2, -3, 4, 5}
	actual := IntsFromString(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
