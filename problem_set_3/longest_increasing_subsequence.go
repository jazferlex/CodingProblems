package main

import "fmt"
func lengthOfLIS(nums []int) int {
	result := []int{nums[0]}
	for _, num := range nums[1:] {
		if num > result[len(result)-1] {
			result = append(result, num)
		} else {
			for i, r := range result {
				if num <= r {
					result[i] = num
					break
				}
			}
		}
	}
	return len(result)
}
func main() {
	test_cases := [][]int{
		{10, 9, 2, 5, 3, 7, 101, 18},
		{0, 1, 0, 3, 2, 3},
		{7, 7, 7, 7, 7, 7, 7},
		{4, 10, 4, 3, 8, 9},
		{2, 2},
		{1, 3, 6, 7, 9, 4, 10, 5, 6},
	}
	for _, test_case := range test_cases {
		fmt.Println(lengthOfLIS(test_case))
	}
}
