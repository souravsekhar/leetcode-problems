package main

import "fmt"

func twoSum(nums []int, target int) []int {
	output := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				output = append(output, i, j)
			}
		}
	}
	return output
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
