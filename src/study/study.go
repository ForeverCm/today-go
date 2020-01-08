package study

import (
	"fmt"
)

func ChanTest() {
	timeout := make (chan bool, 1)
	//go func() {
	//	time.Sleep(1e9) // sleep one second
	//	timeout <- true
	//}()
	//timeout <- true
	ch := make (chan int)
	select {
	case <- ch:
	case <- timeout:
		fmt.Println("timeout!")
	default:
		fmt.Print("default")
	}

}