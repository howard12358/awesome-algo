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

func TestPermuteUnique(t *testing.T) {
	// 定义测试用例结构体，包含输入和期望输出
	type testCase struct {
		input    []int
		expected [][]int
	}

	// 定义测试用例切片
	testCases := []testCase{
		{
			input:    []int{1, 1, 2},
			expected: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			input:    []int{1, 2, 3},
			expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}

	// 遍历每个测试用例
	for _, tc := range testCases {
		// 调用 permuteUnique 函数得到实际输出
		result := permuteUnique(tc.input)

		// 验证实际输出和期望输出是否相等
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("permuteUnique(%v) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "Test1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "Test2",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "Test3",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.candidates, tt.target)
			if len(got) != len(tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if len(got[i]) != len(tt.want[i]) {
					t.Errorf("combinationSum() = %v, want %v", got, tt.want)
					return
				}
				for j := range got[i] {
					if got[i][j] != tt.want[i][j] {
						t.Errorf("combinationSum() = %v, want %v", got, tt.want)
						return
					}
				}
			}
		})
	}
}
