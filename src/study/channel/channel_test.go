package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_CP_ChapterInfo(t *testing.T) {
	fmt.Println("test,test")
	fmt.Println("222")
}

func Test_Select_Timeout(t *testing.T) {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "result 1"
	}()
	for {
		select {
		case res := <-c1:
			fmt.Println(res)
		case <-time.After(time.Second * 2):
			fmt.Println("timeout 1")
			return
		}
	}
	fmt.Println("end............")

}

