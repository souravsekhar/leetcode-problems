package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	index := search(arr, 7)
	fmt.Println(index)
}

func search(nums []int, target int) int {
	l := 0
	h := len(nums) - 1
	for l <= h {
		m := (l + h) / 2
		if nums[m] > target {
			h = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}
