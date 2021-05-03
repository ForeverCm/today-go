package main

import "fmt"

// 贪心+回溯解决硬币找零问题
func main() {
	fmt.Println("hello world")
	fmt.Println(getMainCoinsCount(11, []int{4, 3}))
}

// 硬币找零问题
func getMainCoinsCount(total int, values []int) int {
	if len(values) <= 0 {
		return -1
	}
	currentCoin := values[0]
	currentCount := total / currentCoin
	surplusTotal := total - currentCount * currentCoin
	if surplusTotal == 0 {
		return currentCount
	}

	otherCount := -1
	surplusValues := values[1:]
	for currentCount >= 0 {
		otherCount = getMainCoinsCount(surplusTotal, surplusValues)
		if otherCount == -1 {
			currentCount--
			surplusTotal = total - currentCount * currentCoin
		} else {
			return currentCount + otherCount
		}
	}
	return -1
}


