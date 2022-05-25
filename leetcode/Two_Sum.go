package main

func twoSum(nums []int, target int) []int {
	m := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				m = []int{i, j}
				return m
			}
		}
	}
	return m
}
