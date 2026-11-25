package array

import "sort"

// Problem: 56. Merge Intervals
// Difficulty: Medium
// Link: https://leetcode.com/problems/merge-intervals/
//
// Time Complexity: O(n log n) - Chi phí sắp xếp mảng intervals theo điểm bắt đầu.
// Space Complexity: O(n) - Lưu trữ mảng kết quả (hoặc O(log n) cho stack sắp xếp nếu không tính output).

// merge gộp tất cả các khoảng (intervals) giao nhau và trả về mảng các khoảng không giao nhau.
//
// Ý tưởng giải thuật:
// 1. Sắp xếp các khoảng tăng dần theo giá trị bắt đầu `interval[0]`.
// 2. Khởi tạo mảng `merged` với khoảng đầu tiên.
// 3. Duyệt qua từng khoảng tiếp theo `curr`:
//    - Lấy khoảng cuối cùng trong `merged` là `last`.
//    - Nếu `curr[0] <= last[1]`: Hai khoảng giao nhau, cập nhật `last[1] = max(last[1], curr[1])`.
//    - Ngược lại: Không giao nhau, thêm bản sao của `curr` vào `merged`.
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{{intervals[0][0], intervals[0][1]}}

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		last := merged[len(merged)-1]

		if curr[0] <= last[1] {
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			merged = append(merged, []int{curr[0], curr[1]})
		}
	}

	return merged
}
