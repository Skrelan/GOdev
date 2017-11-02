package arrays

import (
	"fmt"
	"sort"

	"github.com/skrelan/LogrusWrapper/log"
)

func mergeTwoSortedArrays(arr1, arr2 []int, isDescending bool) (*[]int, error) {
	var result = make([]int, 0, len(arr1)+len(arr2))
	var p1, p2 int
	if len(arr1) <= 0 {
		err := fmt.Errorf("len of arr1 less than 0")
		return nil, err
	}
	if len(arr2) <= 0 {
		err := fmt.Errorf("len of arr2 less than 0")
		return nil, err
	}
	for _ = range result {
		if arr1[p1] < arr2[p2] {
			result = append(result, arr1[p1])
			p1++
			if p1 == len(arr1) {
				result = append(result, arr2...)
				break
			}
		} else {
			result = append(result, arr2[p2])
			p2++
			if p2 == len(arr2) {
				result = append(result, arr1...)
				break
			}
		}
	}
	if isDescending {
		sort.Sort(sort.Reverse(sort.IntSlice(result)))
	}
	log.Debug(result)
	return &result, nil
}
