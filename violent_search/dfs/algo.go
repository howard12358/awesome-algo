package dfs

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// 全排列 https://leetcode.cn/problems/permutations/description/
func permute(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(track) == len(nums) {
			temp := append([]int{}, track...)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// 排除重复的选择
			if used[i] {
				// 剪枝，避免重复使用同一个数字
				continue
			}
			// 做选择
			track = append(track, nums[i])
			used[i] = true
			// 进入下一层决策树
			dfs()
			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}
	}

	dfs()
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

// 组合总数 https://leetcode.cn/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var track []int
	backtrack3(candidates, track, target, 0, &res)
	return res
}

// start 是每轮循环中的起始位置，避免重复前面的选择
func backtrack3(candidates []int, track []int, target int, start int, res *[][]int) {
	if target == 0 {
		temp := append([]int{}, track...)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(candidates); i++ {
		// 因为candidates已经排序过了，所以后面的元素都比前面的大，可以直接剪枝
		if target-candidates[i] < 0 {
			return
		}
		track = append(track, candidates[i])
		backtrack3(candidates, track, target-candidates[i], i, res)
		track = track[:len(track)-1]
	}
}

// 组合总和II https://leetcode.cn/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var track []int
	backtrack4(candidates, track, target, 0, &res)
	return res
}

func backtrack4(candidates []int, track []int, target int, start int, res *[][]int) {
	if target == 0 {
		temp := append([]int{}, track...)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			return
		}
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		track = append(track, candidates[i])
		backtrack4(candidates, track, target-candidates[i], i+1, res)
		track = track[:len(track)-1]
	}
}

// 电话号码的字母组合 https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var mapping = [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var res []string
	var track []byte

	var dfs func(i int)
	dfs = func(i int) {
		if i == len(digits) {
			res = append(res, string(track))
			return
		}
		for _, v := range mapping[digits[i]-'0'] {
			track = append(track, byte(v))
			dfs(i + 1)
			track = track[:len(track)-1]
		}
	}
	dfs(0)
	return res
}

// 复原 IP 地址 https://leetcode.cn/problems/restore-ip-addresses/
func restoreIpAddresses(s string) []string {
	var res []string
	var track []string

	var dfs func(start int)
	dfs = func(start int) {
		if len(track) == 4 {
			if start == len(s) {
				res = append(res, strings.Join(track, "."))
			}
			return
		}

		for i := start; i < len(s); i++ {
			if i != start && s[start] == '0' {
				break
			}
			str := s[start : i+1]
			num, _ := strconv.Atoi(str)
			if num >= 0 && num <= 255 {
				track = append(track, str)
				dfs(i + 1)
				track = track[:len(track)-1]
			}
		}
	}
	dfs(0)
	return res
}

// 字母大小写全排列 https://leetcode.cn/problems/letter-case-permutation/description/
func letterCasePermutation(s string) []string {
	var res []string
	var track = []byte(s)
	var dfs func(i int)
	dfs = func(i int) {
		res = append(res, string(track))
		if i == len(s) {
			return
		}
		for j := i; j < len(s); j++ {
			// 如果是字母，就大小写转换
			if unicode.IsLower(rune(track[j])) || unicode.IsUpper(rune(track[j])) {
				track[j] ^= 32
				dfs(j + 1)
				track[j] ^= 32
			}
		}
	}
	dfs(0)
	return res
}

// 分割回文串 https://leetcode.cn/problems/palindrome-partitioning/description/
func partition(s string) [][]string {
	var res [][]string
	var track []string
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			temp := append([]string{}, track...)
			res = append(res, temp)
			return
		}
		for i := start; i < len(s); i++ {
			if isPalindrome(s, start, i) {
				track = append(track, s[start:i+1])
				dfs(i + 1)
				track = track[:len(track)-1]
			}
		}
	}
	dfs(0)
	return res
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
