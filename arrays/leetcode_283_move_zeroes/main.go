package main

import "fmt"

func MoveZeroes(nums []int) {
	lZ := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[lZ] = nums[lZ], nums[i]
			lZ++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums)
	fmt.Println(nums)
}
