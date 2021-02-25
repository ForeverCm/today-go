package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}, 6))
}



//todo 未理解
func subarraySum(nums []int, k int) int {
	sums := make(map[int]int)
	res := 0
	sum := 0
	sums[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res += sums[sum-k]
		sums[sum]++
	}
	return res
}
