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
func maxSum(nums []int, m int, k int) (ans int64) {
	sum := int64(0)
	cntMap := map[int]int{}
	for i, num := range nums {
		sum += int64(num)
		cntMap[num]++
		if i-k+1 < 0 {
			continue
		}

		if len(cntMap) >= m {
			ans = max(ans, sum)
		}

		out := nums[i-k+1]
		sum -= int64(out)
		cntMap[out]--
		if cntMap[out] == 0 {
			delete(cntMap, out)
		}
	}
	return
}

// https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/
func maximumSubarraySum(nums []int, k int) (ans int64) {
	sum := int64(0)
	cntMap := map[int]int{}
	for i, num := range nums {
		sum += int64(num)
		cntMap[num]++
		if i-k+1 < 0 {
			continue
		}

		if len(cntMap) == k {
			ans = max(ans, sum)
		}

		out := nums[i-k+1]
		sum -= int64(out)
		cntMap[out]--
		if cntMap[out] == 0 {
			delete(cntMap, out)
		}
	}
	return
}

// https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards/description/
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	sum := 0
	for _, num := range cardPoints {
		sum += num
	}
	if k == n {
		return sum
	}
	s := 0
	minS := sum
	for i, num := range cardPoints {
		s += num
		left := i - (n - k) + 1
		if left < 0 {
			continue
		}
		minS = min(minS, s)

		out := cardPoints[left]
		s -= out
	}
	return sum - minS
}

// https://leetcode.cn/problems/grumpy-bookstore-owner/
func maxSatisfied(customers []int, grumpy []int, minutes int) (ans int) {
	m := map[int]int{
		0: 0,
		1: 0,
	}
	maxC := 0
	for i, num := range customers {
		m[grumpy[i]] += num
		if i-minutes+1 < 0 {
			continue
		}
		maxC = max(maxC, m[1])
		if grumpy[i-minutes+1] == 1 {
			m[1] -= customers[i-minutes+1]
		}
	}
	return m[0] + maxC
}
