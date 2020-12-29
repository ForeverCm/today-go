package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 1, 3, 4}))
}


//todo 未理解
func findUnsortedSubarray(n []int) int {
	l := len(n)
	for i := 1; i < len(n); i++ {
		if n[i] < n[i-1] {
			l = i-1
			break
		}
	}
	if l == len(n) {
		return 0
	}
	r := -1
	for i := len(n) - 2; i >= 0; i-- {
		if n[i] > n[i+1] {
			r = i + 1
			break
		}
	}
	min, max := n[l], n[r]
	for i := l; i <= r; i++ {
		if n[i] < min {
			min = n[i]
		} else if n[i] > max {
			max = n[i]
		}
	}
	l, r = -1, -1
	for i := 0; i < len(n); i++ {
		if n[i] > min && l == -1 {
			l = i
		}
		if n[i] < max {
			r = i
		}
	}
	return r - l + 1
}

func findUnsortedSubarray3(nums []int) int {
	dst := make([]int, len(nums))
	copy(dst, nums)
	sort.Ints(dst)
	l := -1
	r := -1
	for i := 0; i < len(dst); i++ {
		if dst[i] != nums[i] {
			if l == -1 {
				l = i
			} else {
				r = i
			}
		}
	}
	if r == -1 {
		return 0
	}
	return r - l + 1
}








func findUnsortedSubarray2(nums []int) int {
	dst := make([]int, len(nums))
	copy(dst, nums)
	sort.Ints(nums)
	var l, r int = -1, -1
	for i := range nums {
		if nums[i] != dst[i] {
			if l == - 1 {
				l = i
			} else {
				r = i
			}
		}
	}
	if r == -1 {
		return 0
	}
	return r - l + 1
}
