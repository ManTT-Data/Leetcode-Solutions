package greedy

// Problem: 134. Gas Station
// Difficulty: Medium
// Link: https://leetcode.com/problems/gas-station/
//
// Time Complexity: O(n) - Duyệt qua mảng đúng một lần.
// Space Complexity: O(1) - Chỉ sử dụng các biến đếm tổng và chỉ số xuất phát.

// canCompleteCircuit tìm chỉ số trạm xăng bắt đầu để có thể đi một vòng tròn mà không bị hết xăng giữa đường.
// Sử dụng thuật toán Tham lam (Greedy Algorithm):
// 1. Nhận xét tổng quát:
//    - Nếu tổng lượng xăng (total gas) < tổng chi phí (total cost), chắc chắn không thể hoàn thành một vòng -> trả về -1.
//    - Ngược lại, nếu tổng lượng xăng >= tổng chi phí, định lý đảm bảo luôn tồn tại một điểm xuất phát hợp lệ duy nhất.
//
// 2. Lựa chọn tham lam (Greedy Choice):
//    - Giả sử xuất phát từ trạm `start`, ta duy trì lượng xăng hiện tại `currentGas`.
//    - Khi đi tới trạm `i` mà `currentGas < 0` (hết xăng dọc đường):
//      Điều này chứng minh mọi trạm nằm giữa `start` và `i` đều KHÔNG THỂ là điểm xuất phát hợp lệ
//      (bởi vì xuất phát từ `start` ta đến các trạm đó với lượng xăng >= 0 mà còn bị hết xăng tại `i`,
//      thì nếu xuất phát từ một trạm ở giữa với bình xăng rỗng = 0 chắc chắn sẽ chết máy tại hoặc trước `i`).
//    - Do đó, ta tham lam dời điểm xuất phát sang trạm kế tiếp: `start = i + 1`, và reset `currentGas = 0`.
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	currentGas := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		totalGas += diff
		currentGas += diff

		// Nếu hết xăng tại trạm i, trạm xuất phát không thể nằm trong khoảng [start, i]
		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
	}

	// Nếu tổng lượng xăng tích lũy toàn bộ chuyến đi âm, không thể hoàn thành vòng đua
	if totalGas < 0 {
		return -1
	}

	return start
}
