package main

// TODO Revisar
import "fmt"

func twoSums(nums []int, target int) []int {
	res := make([]int, 2)
	resMap := make(map[int]int)

	for i, num := range nums {
		c := target - num
		if index, ok := resMap[c]; ok {
			res = []int{index, i}
			break
		}
		resMap[num] = i
	}
	return res
}

func main() {
	fmt.Println(twoSums([]int{2, 7, 11, 15}, 9))
}
