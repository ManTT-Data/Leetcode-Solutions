package greedy

// Problem: 4026. Maximum Gap Between Stations
// Difficulty: Medium
// Link: https://leetcode.com/problems/maximum-gap-between-stations/
//
// Time Complexity: O(n + m) - Duyệt qua chuỗi station hai lần (từ trái sang phải và từ phải sang trái).
// Space Complexity: O(n) - Lưu trữ hai mảng vị trí khớp leftMatch và rightMatch kích thước n.

// maxGapBetweenStations tìm khoảng cách lớn nhất giữa hai trạm gán cho hai công nhân liên tiếp.
//
// Ý tưởng giải thuật: Two-way Greedy Matching (Khớp tham lam hai chiều):
// 1. Nhận xét:
//    - Để tối đa hóa khoảng cách giữa hai công nhân liên tiếp (i - 1) và i (tức j_i - j_{i-1}):
//      + Ta muốn công nhân (i - 1) được gán vào trạm CÀNG NHỎ (càng về bên trái) càng tốt.
//      + Ta muốn công nhân i được gán vào trạm CÀNG LỚN (càng về bên phải) càng tốt.
// 2. Xây dựng hai mảng tham lam:
//    - `leftMatch[k]`: Vị trí nhỏ nhất trong `station` có thể gán cho `skill[k]` sao cho toàn bộ
//      tiền tố `skill[0..k]` được khớp hợp lệ (duyệt tham lam từ trái sang phải).
//    - `rightMatch[k]`: Vị trí lớn nhất trong `station` có thể gán cho `skill[k]` sao cho toàn bộ
//      hậu tố `skill[k..n-1]` được khớp hợp lệ (duyệt tham lam từ phải sang trái).
// 3. Với mỗi ranh giới i (1 <= i < n):
//    - Khoảng cách lớn nhất giữa công nhân (i - 1) và i là: rightMatch[i] - leftMatch[i-1].
//    - Kết quả là giá trị cực đại trên tất cả các ranh giới i.
//    - Nếu chỉ có 1 công nhân (n == 1), khoảng cách mặc định là 0.
func maxGapBetweenStations(skill string, station string) int {
	n := len(skill)
	m := len(station)

	if n <= 1 {
		return 0
	}

	leftMatch := make([]int, n)
	rightMatch := make([]int, n)

	// Khớp tham lam từ trái sang phải
	p := 0
	for i := 0; i < n; i++ {
		for p < m && station[p] != skill[i] {
			p++
		}
		leftMatch[i] = p
		p++
	}

	// Khớp tham lam từ phải sang trái
	p = m - 1
	for i := n - 1; i >= 0; i-- {
		for p >= 0 && station[p] != skill[i] {
			p--
		}
		rightMatch[i] = p
		p--
	}

	// Tìm khoảng cách lớn nhất giữa hai công nhân liên tiếp
	maxGap := 0
	for i := 1; i < n; i++ {
		gap := rightMatch[i] - leftMatch[i-1]
		if gap > maxGap {
			maxGap = gap
		}
	}

	return maxGap
}
