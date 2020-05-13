package main

import (
	"fmt"
	"golang.org/x/cryptov0.0.0-20200210222208-86ce3cb69678/src/leetcode"
	"math"
	"net"
)

func main() {

	fmt.Println(leetcode.IsMatchNew("aaa", "bbb"))
	//PushTask()
}

func PushTask()  {
	defer func() {
		if err := recover(); err != nil{
			go PushTask()
		}
	}()
	for {
		pushHandle()
	}
}

func pushHandle() {
	//fmt.Println("hello world")
	panic("test")
}


func get_internal() string{
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}


func printPrime(prefix string) {

		for outer := 2; outer < 5000; outer++ {
		next:
			for inner := 2; inner < outer; inner++ {
				if outer % inner == 0 {
					continue next
				}
			}
			fmt.Printf("%s:%d\n", prefix, outer)
		}
	fmt.Println("Completed", prefix)
}





func myAtoi(str string) int {
	var (
		ans   int64
		start bool
		sign  = 1
	)

	for _, v := range str {
		if !start && v == ' ' {
			continue
		}
		if '0' <= v && v <= '9' {
			start = true
			ans = ans * 10 + int64(v - '0')
			if ans > math.MaxInt32+1 {
				break
			}
		} else if !start && v == '+' {
			start = true
			sign = 1
		} else if !start && v == '-' {
			start = true
			sign = -1
		} else {
			break
		}
	}

	ans *= int64(sign)
	if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	if ans < math.MinInt32 {
		return math.MinInt32
	}

	return int(ans)
}