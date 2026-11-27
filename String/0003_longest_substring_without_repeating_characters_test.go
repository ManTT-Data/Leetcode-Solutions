package string_problems

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1: abcabcbb",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "Example 2: bbbbb",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "Example 3: pwwkew",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "Empty string",
			s:    "",
			want: 0,
		},
		{
			name: "Single space",
			s:    " ",
			want: 1,
		},
		{
			name: "abba (left pointer must not move backward)",
			s:    "abba",
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d; want %d", tt.s, got, tt.want)
			}
		})
	}
}
