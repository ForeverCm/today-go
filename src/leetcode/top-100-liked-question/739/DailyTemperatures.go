package main

import "fmt"

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}

func dailyTemperatures(T []int) []int {
	if len(T) <= 0 {
		return []int{}
	}
	var res []int
	for i := 0; i < len(T); i++ {
		j := i+1
		for ;j < len(T); j++ {
			if T[j] > T[i] {
				res = append(res, j-i)
				break
			}
		}
		if j == len(T) {
			res = append(res, 0)
		}
	}
	return res
}
