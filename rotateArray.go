package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
}

func rotate(nums []int, k int) {

	if len(nums) < k {
		k = (k % len(nums))
	}

	nums = reverse(nums, 0, len(nums)-k-1)
	nums = reverse(nums, len(nums)-k, len(nums)-1)
	nums = reverse(nums, 0, len(nums)-1)

	fmt.Println(nums)
}

func reverse(reversed []int, s int, e int) []int {

	temp := 0
	for s <= e {
		temp = reversed[s]
		reversed[s] = reversed[e]
		reversed[e] = temp

		s++
		e--
	}
	return reversed
}
