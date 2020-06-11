package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// for ; i < 10; i++ {
	// 	i := i
	// 	go func() {
	// 		fmt.Println(i)
	// 	}()
	// }

	// time.Sleep(2 * time.Second)

	l := []int{1, 2, 3}
	for idx, item := range l {
		theIdx, theItem := idx, item
		go func() {
			fmt.Println(theIdx, theItem)
		}()
	}
	time.Sleep(time.Second)

	// l := []int{1, 2, 3}
	// for idx, item := range l {
	// 	go func() {
	// 		fmt.Println(idx, item)
	// 	}()
	// }
	// time.Sleep(10 * time.Second)
}
