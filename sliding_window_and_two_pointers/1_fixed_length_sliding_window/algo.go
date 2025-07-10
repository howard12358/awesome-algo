package __fixed_length_sliding_window

import (
	"math"
)

// https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/
func maxVowels(s string, k int) (ans int) {
	vowel := 0
	for i, in := range s {
		if in == 'a' || in == 'e' || in == 'i' || in == 'o' || in == 'u' {
			vowel++
		}
		if i-k+1 < 0 {
			continue
		}
		ans = max(vowel, ans)
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			vowel--
		}
	}
	return
}

// https://leetcode.cn/problems/maximum-average-subarray-i/
func findMaxAverage(nums []int, k int) (ans float64) {
	maxS := math.MinInt
	sum := 0
	for i, num := range nums {
		sum += num
		if i-k+1 < 0 {
			continue
		}
		maxS = max(sum, maxS)
		sum -= nums[i-k+1]
	}
	return float64(maxS) / float64(k)
}

// https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
func numOfSubarrays(arr []int, k int, threshold int) (ans int) {
	sum := 0
	for i, num := range arr {
		sum += num
		if i-k+1 < 0 {
			continue
		}
		if sum/k >= threshold {
			ans++
		}
		sum -= arr[i-k+1]
	}
	return
}

// https://leetcode.cn/problems/k-radius-subarray-averages/description/
func getAverages(nums []int, k int) (ans []int) {
	ans = make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	sum := 0
	for i, num := range nums {
		sum += num
		if i-2*k < 0 {
			continue
		}
		ans[i-k] = sum / (2*k + 1)
		sum -= nums[i-2*k]
	}
	return
}

// https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/description/
func minimumRecolors(blocks string, k int) int {
	countW := 0
	ans := math.MaxInt
	for i, ch := range blocks {
		if ch == 'W' {
			countW++
		}
		if i-k+1 < 0 {
			continue
		}
		ans = min(ans, countW)
		if blocks[i-k+1] == 'W' {
			countW--
		}
	}
	return ans
}

// https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray/description/
func maxSum(nums []int, m int, k int) int64 {

	return 0
}
