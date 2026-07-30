/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30404
 *
 * [1] 两数之和
 */

package leetcode_solutions

import "testing"

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := hashMap[complement]; ok {
			return []int{j, i}
		}
		hashMap[num] = i
	}
	return nil
}

// @lc code=end

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"case1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"case2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"case3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if len(got) != len(tt.want) || got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/
