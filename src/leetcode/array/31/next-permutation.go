package main

import "fmt"

func main() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
func nextPermutation(nums []int)  {
	var j, i int
	i = len(nums) - 2

	for i >= 0 {
		if nums[i] < nums[i+1] {

			j = i + 1
			for j < len(nums) && nums[j] > nums[i] {
				j++
			}
			j--
			nums[i], nums[j] = nums[j], nums[i]
			break
		}
		i--
	}
	i++
	for j = len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}