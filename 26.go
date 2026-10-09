package main

import (
	"slices"
)

func removeDuplicates(nums []int) int {
	slow := 0
	fast := 1
	for fast <= len(nums)-1 {

		if nums[slow] != nums[fast] {
			slow++
			fast++
		} else {
			nums = slices.Delete(nums, fast, fast+1)
		}
	}
	return len(nums)
}

// func main() {
// 	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
// }
