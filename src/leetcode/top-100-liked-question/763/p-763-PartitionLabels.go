package main

import "fmt"

func main() {
	fmt.Println("aaabc")
	fmt.Println(partitionLabels("aaabcaddf"))
}

func partitionLabels(S string) []int {
	if len(S) <= 0 {
		return []int{}
	}
	temp := map[byte]int{}
	for i := 0; i < len(S); i++ {
		temp[S[i]] = i
	}
	start := 0
	end := temp[S[start]]
	var res []int
	for i := 0; i < len(S); i++ {
		if i == end {
			res = append(res, end - start + 1)
			start = i + 1
			if start == len(S) {
				return res
			}
			end = temp[S[start]]
		} else {
			end = max(end, temp[S[i]])
		}
	}
	return res

}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}