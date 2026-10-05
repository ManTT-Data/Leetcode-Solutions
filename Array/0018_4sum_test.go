package array

import (
	"reflect"
	"sort"
	"testing"
)

func canonicalizeQuadruplets(quads [][]int) [][]int {
	res := make([][]int, len(quads))
	for i, q := range quads {
		cp := make([]int, len(q))
		copy(cp, q)
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

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{
			name:   "Example 1",
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			name:   "Example 2",
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want: [][]int{
				{2, 2, 2, 2},
			},
		},
		{
			name:   "Less than 4 elements",
			nums:   []int{1, 2, 3},
			target: 6,
			want:   nil,
		},
		{
			name:   "Potential 32-bit overflow test",
			nums:   []int{1000000000, 1000000000, 1000000000, 1000000000},
			target: -294967296,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.nums, tt.target)
			cGot := canonicalizeQuadruplets(got)
			cWant := canonicalizeQuadruplets(tt.want)

			if len(cGot) == 0 && len(cWant) == 0 {
				return
			}

			if !reflect.DeepEqual(cGot, cWant) {
				t.Errorf("fourSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
