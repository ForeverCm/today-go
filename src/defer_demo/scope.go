package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println(x)
	x = 88
	fmt.Println(x)
	{
		x = 56
		fmt.Println(x)
		y := "The test message"
		fmt.Println(y)
	}
}
