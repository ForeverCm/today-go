package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDisappearedNumbers([]int{4,3,2,7,8,2,3,1}))
}

func findDisappearedNumbers(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}
	length := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 1 || nums[i] > length {
			return nil
		}
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	var result []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			result = append(result, i+1)
		}
	}
	return result
}