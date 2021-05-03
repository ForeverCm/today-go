package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(fibonacci(10))
	time1 := time.Now()
	fmt.Println(fibonacci2(20))
	fmt.Println(time.Now().Sub(time1).Nanoseconds())
	time2 := time.Now()
	memo := make([]int, 100)
	fmt.Println(fibonacciAdvanced(20, memo))
	fmt.Println(time.Now().Sub(time2).Nanoseconds())

}

func fibonacci(n int) int {
	if n < 0 {
		return -1
	}
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return -1
	}
	if n < 2 {
		return n
	}
	first := 0
	second := 1
	third := 1
	for n > 1 {
		third = first + second
		first = second
		second = third
		n--
	}
	return third
}

func fibonacciAdvanced(n int, memo []int) int {
	if n < 2 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	if n >= 2 {
		memo[n] = fibonacciAdvanced(n-1, memo) + fibonacciAdvanced(n-2, memo)
		return memo[n]
	}
	return -1
}