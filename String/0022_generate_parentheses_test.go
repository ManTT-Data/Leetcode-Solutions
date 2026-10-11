package string_problems

import (
	"reflect"
	"sort"
	"testing"
)

func isValidParentheses(s string) bool {
	balance := 0
	for _, ch := range s {
		if ch == '(' {
			balance++
		} else {
			balance--
			if balance < 0 {
				return false
			}
		}
	}
	return balance == 0
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "Example 2: n = 1",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "n = 2",
			n:    2,
			want: []string{"(())", "()()"},
		},
		{
			name: "Example 1: n = 3",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)

			gotSorted := append([]string(nil), got...)
			wantSorted := append([]string(nil), tt.want...)
			sort.Strings(gotSorted)
			sort.Strings(wantSorted)

			if !reflect.DeepEqual(gotSorted, wantSorted) {
				t.Errorf("generateParenthesis(%d) = %v; want %v", tt.n, got, tt.want)
			}

			// Kiểm tra tất cả chuỗi sinh ra đều hợp lệ
			for _, s := range got {
				if !isValidParentheses(s) {
					t.Errorf("generateParenthesis(%d) generated invalid parentheses: %s", tt.n, s)
				}
			}
		})
	}
}

func TestGenerateParenthesisCatalanCount(t *testing.T) {
	// Số Catalan thứ n: C_1=1, C_2=2, C_3=5, C_4=14, C_5=42
	catalanCounts := map[int]int{
		1: 1,
		2: 2,
		3: 5,
		4: 14,
		5: 42,
	}

	for n, expectedCount := range catalanCounts {
		got := generateParenthesis(n)
		if len(got) != expectedCount {
			t.Errorf("generateParenthesis(%d) returned %d items; want %d", n, len(got), expectedCount)
		}

		// Đảm bảo không có chuỗi trùng lặp
		seen := make(map[string]bool)
		for _, s := range got {
			if seen[s] {
				t.Errorf("duplicate parentheses %s found for n = %d", s, n)
			}
			seen[s] = true
			if !isValidParentheses(s) {
				t.Errorf("invalid parentheses %s found for n = %d", s, n)
			}
		}
	}
}
