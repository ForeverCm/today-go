package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstrings("aba"))
	fmt.Println(countSubstrings("aaaa"))
}

func countSubstrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	res := 0
	for i := 1; i <= len(s); i++ {
		start := 0
		end := start + i
		for end <= len(s) {
			if isPalindromic(s[start : end]) {
				res++
			}
			start += 1
			end = start + i
		}
	}
	return res
}


func isPalindromic(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	for i := 0; i < len(s) / 2; i++ {
		if s[i] != s[len(s) - i - 1] {
			return false
		}
	}
	return true
}
