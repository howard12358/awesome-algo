package fixed_length_sliding_window

import "math"

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
func numOfSubarrays(arr []int, k int, threshold int) int {

}
