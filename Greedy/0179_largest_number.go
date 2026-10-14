package greedy

import (
	"sort"
	"strconv"
	"strings"
)

// Problem: 179. Largest Number
// Difficulty: Medium
// Link: https://leetcode.com/problems/largest-number/
//
// Time Complexity: O(n log n * k) - Trong đó n là số phần tử, k là độ dài tối đa của số dạng chuỗi (k <= 10). Phép so sánh chuỗi a + b tốn O(k).
// Space Complexity: O(n * k) - Lưu trữ mảng chuỗi biểu diễn các số.

// largestNumber sắp xếp các số nguyên không âm sao cho ghép lại thành số lớn nhất.
// Sử dụng thuật toán Tham lam thông qua Sắp xếp tùy biến (Greedy Custom Sorting):
// - Với hai chuỗi số a và b, ta so sánh hai cách ghép: (a + b) và (b + a).
//   + Nếu (a + b) > (b + a): a nên đứng trước b.
//   + Nếu (a + b) < (b + a): b nên đứng trước a.
// - Tính chất bắc cầu (Transitivity): Quy tắc so sánh này đảm bảo thứ tự sắp xếp toàn cục
//   dẫn đến số lớn nhất có thể tạo thành.
// - Trường hợp đặc biệt: Nếu sau khi sắp xếp, số lớn nhất ở vị trí đầu tiên là "0"
//   (tức là mảng chỉ toàn số 0 như [0, 0]), ta trả về "0" thay vì "00".
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	// Sắp xếp giảm dần theo quy tắc ghép chuỗi tham lam
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	// Nếu số lớn nhất là "0", toàn bộ mảng là các số 0
	if strs[0] == "0" {
		return "0"
	}

	return strings.Join(strs, "")
}
