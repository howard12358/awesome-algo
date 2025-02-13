package dfs

// 全排列 https://leetcode.cn/problems/permutations/description/
func permute(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))

	backtrack(nums, track, used, &res)
	return res
}

func backtrack(nums []int, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := range nums {
		if used[i] {
			continue
		}
		// 做选择
		track = append(track, nums[i])
		used[i] = true
		// 进入下一层决策树
		backtrack(nums, track, used, res)
		// 撤销选择
		track = track[:len(track)-1]
		used[i] = false
	}
}
