package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	//println(runtime.NumCPU())
	//runtime.GOMAXPROCS(4)
	for i := 0; i < 50; i = i + 1 {
		//wg.Add(1)
		//fmt.Println(strconv.Itoa(i))
		go func(n int) {
			//defer wg.Done()
			//defer wg.Add(-1)
			fmt.Println(strconv.Itoa(i) + ":::" + strconv.Itoa(n))
		}(i)

		//runtime.Gosched()
		//time.Sleep(time.Millisecond)
	}
	//wg.Wait()
	//println(runtime.NumCPU())
	time.Sleep(time.Second)

	//ch := make(chan int, 1)
	//for i := 0; i < 10; i++ {
	//	select {
	//	case x := <-ch:
	//
	//		fmt.Println(x) // "0" "2" "4" "6" "8"
	//	case ch <- i:
	//
	//	}
	//}

}
