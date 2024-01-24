package main

import "fmt"

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1
	for i := 0; i < len(list); i++ {
		mid := (low + high) / 2
		if list[mid] == item {
			return mid
		} else if list[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 40, 50, 60, 75, 100, 102}, 24))
	fmt.Println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
}
