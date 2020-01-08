package leetcode

import "strings"

func isMatch(s string, p string) bool {

	m, n := len(s), len(p)
	dp := make([][]bool, m + 1)
	for i := range dp {
		dp[i] = make([]bool, n + 1)
	}

	//init
	//s and p are empty
	dp[0][0] = true
	//s is empty
	for j := 1; j <= n; j++ {
		if p[j - 1] == '.' {
			dp[0][j] = false
		} else if p[j - 1] == '*' {
			dp[0][j] = dp[0][j-1] || (j >= 2 && dp[0][j-2])
		} else {
			dp[0][j] = false
		}
	}

	//calculate
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = (j > 1 && dp[i][j-2]) || (dp[i][j-1]) ||
					(dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			} else {
				dp[i][j] = (s[i-1]==p[j-1]) && (dp[i-1][j-1])
			}
		}
	}

	//result
	return dp[m][n]
}


func IsMatchNew(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	if len(p) == 1 {
		return len(s) == 1 && (p[0] == s[0] || p[0] == '.')
	}

	if p[1] != '*' {
		if s == "" {
			return false
		}

		return (p[0] == s[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}

	if s != "" && (s[0] == p[0] || p[0] == '.') {
		if IsMatchNew(s, p[2:]) {
			return true
		}
		s = s[1:]
	}

	return IsMatchNew(s, p[2:])
}



func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	sum := m + n
	isodd := isOdd(sum)

	var left int
	if (isodd) {
		left = (sum -1) / 2
	} else {
		left = sum / 2
	}

	counter := 0
	i, j := 0, 0
	leftValue, rightValue := 0, 0

	for i < m || j < n {
		leftValue = rightValue
		if i < m && j < n {
			if nums1[i] <= nums2[j] {
				rightValue = nums1[i]
				i += 1
			} else {
				rightValue = nums2[j]
				j += 1
			}
		} else if i < m {
			rightValue = nums1[i]
			i += 1
		} else if j < n {
			rightValue = nums2[j]
			j += 1
		}

		counter += 1

		if counter < sum {
			if isodd && counter - 2 == left {
				break
			}
			if !isodd && counter -1 == left {
				break
			}
		}
	}

 	if isodd {
		return float64(leftValue + rightValue) / 2
	}
	return float64(rightValue + rightValue) / 2
}



func intToRoman(num int) string {
	var b strings.Builder
	m := [13]int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	n := [13]string {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < 13; i++ {
		j := num / m[i]
		num %= m[i]
		for j > 0 {
			b.WriteString(n[i])
			j--
		}
	}
	return b.String()
}


func isOdd(param int) bool {
	if param & 1 == 0 {
		return true
	}
	return false
}
