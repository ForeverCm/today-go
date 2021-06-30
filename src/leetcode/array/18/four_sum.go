package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{-2,-1,-1,1,1,2,2}, 0))
}
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	count := len(nums)
	if count < 4 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < count; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i+1; j < count-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k, m := j+1, count-1
			for k < m {
				if k > j+1 && nums[k] == nums[k-1] {
					k++
				}
				if m < count-1 && nums[m] == nums[m+1] {
					m--
				}
				cur := nums[i] + nums[j] + nums[k] + nums[m]
				if cur == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[m]})
					k++
					m--
				}else if cur < target{
					k++
				}else {
					m--
				}
			}
		}
	}
	return res
}