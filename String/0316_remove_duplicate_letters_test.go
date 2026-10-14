package string_problems

import "testing"

func TestRemoveDuplicateLetters(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "bcabc",
			want: "abc",
		},
		{
			name: "Example 2",
			s:    "cbacdcbc",
			want: "acdb",
		},
		{
			name: "Single character",
			s:    "a",
			want: "a",
		},
		{
			name: "All identical characters",
			s:    "aaaa",
			want: "a",
		},
		{
			name: "Already sorted and distinct",
			s:    "abc",
			want: "abc",
		},
		{
			name: "Strictly decreasing distinct",
			s:    "cba",
			want: "cba",
		},
		{
			name: "Duplicate with later smaller characters",
			s:    "bbcaac",
			want: "bac",
		},
		{
			name: "Complex case abacb",
			s:    "abacb",
			want: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicateLetters(tt.s)
			if got != tt.want {
				t.Errorf("removeDuplicateLetters(%q) = %q; want %q", tt.s, got, tt.want)
			}
		})
	}
}
