package dfs

// 全排列 https://leetcode.cn/problems/permutations/description/
func permute(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))

	backtrack(nums, track, used, &res)
	return res
}

// track 表示 nums 已经选择的路径
// used 表示已经做出的选择和已经剩余的选择
func backtrack(nums []int, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := range nums {
		// 排除重复的选择
		if used[i] {
			// 剪枝，避免重复使用同一个数字
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

// 全排列II https://leetcode.cn/problems/permutations-ii/description/
func permuteUnique(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))
	backtrack2(nums, track, used, &res)
	return res
}

func backtrack2(nums []int, track []int, used []bool, res *[][]int) {
	if len(nums) == len(track) {
		temp := append([]int{}, track...)
		*res = append(*res, temp)
		return
	}
	// 记录已经尝试过的元素，保证本轮循环中相等元素只被选择一次
	// https://www.hello-algo.com/chapter_backtracking/permutations_problem/#3
	duplicated := make(map[int]struct{})
	for i := range nums {
		_, ok := duplicated[nums[i]]
		// 剪枝，nums中相同位置的元素和相同的元素都被剪去
		if used[i] || ok {
			continue
		}
		duplicated[nums[i]] = struct{}{}
		used[i] = true
		track = append(track, nums[i])
		backtrack2(nums, track, used, res)
		track = track[:len(track)-1]
		used[i] = false
	}
}
