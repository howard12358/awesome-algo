package dfs

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	// 定义测试用例
	testCases := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{1, 2, 3},
			expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			input:    []int{0, 1},
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			input:    []int{1},
			expected: [][]int{{1}},
		},
	}

	// 遍历每个测试用例
	for _, tc := range testCases {
		// 调用 permute 函数得到实际结果
		result := permute(tc.input)

		// 检查实际结果和期望结果是否相等
		if !reflect.DeepEqual(result, tc.expected) {
			// 如果不相等，输出错误信息
			t.Errorf("permute(%v) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}
