package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	reverse(arr, 0, len(arr)-1)
}

func reverse(arr []int, start int, end int) {
	temp := 0
	for start <= end {
		temp = arr[start]
		arr[start] = arr[end]
		arr[end] = temp

		start++
		end--
	}

	fmt.Println(arr)
}
