package array

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name: "Example 1 - 3x3 square matrix",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "Example 2 - 3x4 rectangular matrix",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "Single element matrix 1x1",
			matrix: [][]int{
				{1},
			},
			want: []int{1},
		},
		{
			name: "Single row matrix 1x4",
			matrix: [][]int{
				{1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Single column matrix 4x1",
			matrix: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "2x2 square matrix",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			want: []int{1, 2, 4, 3},
		},
		{
			name: "4x3 matrix with negative numbers",
			matrix: [][]int{
				{1, -2, 3},
				{4, 5, -6},
				{-7, 8, 9},
				{10, -11, 12},
			},
			want: []int{1, -2, 3, -6, 9, 12, -11, 10, -7, 4, 5, 8},
		},
		{
			name:   "Empty matrix",
			matrix: [][]int{},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := spiralOrder(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder(%v) = %v; want %v", tt.matrix, got, tt.want)
			}
		})
	}
}
