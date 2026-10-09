package twopointer

import (
	"reflect"
	"sort"
	"testing"
)

// canonicalizeTriplets sorts each triplet and then sorts the list of triplets for consistent comparison.
func canonicalizeTriplets(triplets [][]int) [][]int {
	res := make([][]int, len(triplets))
	for i, t := range triplets {
		cp := make([]int, len(t))
		copy(cp, t)
		sort.Ints(cp)
		res[i] = cp
	}

	sort.Slice(res, func(i, j int) bool {
		for k := 0; k < len(res[i]) && k < len(res[j]); k++ {
			if res[i][k] != res[j][k] {
				return res[i][k] < res[j][k]
			}
		}
		return len(res[i]) < len(res[j])
	})

	return res
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 1},
			want: nil,
		},
		{
			name: "Example 3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Duplicates with positive and negative numbers",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			cGot := canonicalizeTriplets(got)
			cWant := canonicalizeTriplets(tt.want)

			if len(cGot) == 0 && len(cWant) == 0 {
				return
			}

			if !reflect.DeepEqual(cGot, cWant) {
				t.Errorf("threeSum(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}
