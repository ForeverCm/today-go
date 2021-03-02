package main

import (
	"fmt"
	"sync"
)

func main() {
	var numsInt []int

	responseChannel := make(chan int, 10)
//	done := make(chan int)
	go func() {
		for rc := range responseChannel {
			numsInt = append(numsInt, rc)
		}
//		close(done)
	}()

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, c chan int) {
			defer wg.Done()
			data := num
			c <- data
		}(i, responseChannel)
	}
	wg.Wait()

//	close(responseChannel)
//	<-done

	// 这里打印的len(numsInt) < 10
	fmt.Println("len(numsInt) = ", len(numsInt))
	// 这里打印了10个数字
	fmt.Println("numsInt = ", numsInt)
	// 这里打印的len(numsInt) == 10
	fmt.Println("len(numsInt) = ", len(numsInt))
}
