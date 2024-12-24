package algo_test

import (
	"algo"
	"testing"
)

func TestBinarySearchInt(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 7, 9, 11}, 7, 3},
		{[]int{1, 3, 5, 7, 9, 11}, 1, 0},
		{[]int{1, 3, 5, 7, 9, 11}, 11, 5},
	}

	for _, test := range tests {
		result := algo.BinarySearchInt(test.arr, test.target)
		if result != test.expected {
			t.Errorf("For array %v and target %d, expected %d but got %d",
				test.arr, test.target, test.expected, result)
		}
	}
}
