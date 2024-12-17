package main

import (
	"reflect"
	"testing"
)

func Test_executeProgram(t *testing.T) {
	_, _, b, _ := executeProgram([]int{2, 6}, 0, 0, 9)
	if b != 1 {
		t.Errorf("Expected b=1, got %d", b)
	}

	output, _, _, _ := executeProgram([]int{5, 0, 5, 1, 5, 4}, 10, 0, 0)
	if !reflect.DeepEqual(output, []int{0, 1, 2}) {
		t.Errorf("Expected %v, got %v", []int{0, 1, 2}, output)
	}
}
