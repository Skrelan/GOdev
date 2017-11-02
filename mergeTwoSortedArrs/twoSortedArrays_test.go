package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSortedArays(t *testing.T) {
	type testDataStruct struct {
		arr1         []int
		arr2         []int
		isDescending bool
	}
	testData := []testDataStruct{
		{arr1: []int{1, 3, 5, 6, 7}, arr2: []int{-1, 0, 1, 4, 6, 17}, isDescending: false},
		{arr1: []int{1, 3, 5, 6, 7}, arr2: []int{-1, 0, 1, 4, 6, 17}, isDescending: true},
		{arr1: []int{}, arr2: []int{1, 3, 4, 5, 8}},
	}
	for i, data := range testData {
		_, err := mergeTwoSortedArrays(data.arr1, data.arr2, data.isDescending)
		if i < 2 {
			assert.NoError(t, err, "This should not return an error")
		} else {
			assert.Error(t, err, "This should return an error")
		}
	}
}
