package image

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// 等待所有 goroutine 执行完毕
// 使用传址方式为 WaitGroup 变量传参
// 使用 channel 关闭 goroutine

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	ch := make(chan interface{}, 10)
	runtime.GOMAXPROCS(runtime.NumCPU())
	var count = 0
	//workerCount := 400
	runtime.GOMAXPROCS(2)
	dir, err := ioutil.ReadDir("E:/mzitu")
	if err != nil {
		fmt.Println("WWW")
	}
	for _, fi := range dir {
		wg.Add(1)
		fmt.Println(fi.Name())
		count++
		go doIt(count, ch, done, fi, &wg) // wg 传指针，doIt() 内部会改变 wg 的值
	}

	for i := 0; i < 10; i++ { // 向 ch 中发送数据，关闭 goroutine
		ch <- i
	}

	close(done)
	wg.Wait()
	close(ch)
	fmt.Println("all done!")
	//dir, err := ioutil.ReadDir("E:/mzitu")
	//if err != nil {
	//	fmt.Println("WWW")
	//}
	//var count = 0
	//for _, fi := range dir {
	//	count ++
	//	fmt.Println(strconv.Itoa(count) + fi.Name())
	//}
}

func doIt(workerID int, ch <-chan interface{}, done <-chan struct{}, fi os.FileInfo, wg *sync.WaitGroup) {
	time.Sleep(100000)
	//fmt.Printf("[%v] is running  %s \n", workerID, fi.Name())
	PthSep := string(os.PathSeparator)
	newPath := "E:\\mzi" + PthSep + "tuku_" + strconv.Itoa(workerID)
	if fi.IsDir() {
		workerID++
		//fmt.Println("E:/mzitu/" + fi.Name())
		dir, err := ioutil.ReadDir("E:/mzitu/" + fi.Name())
		if err != nil {
			fmt.Println(err)
		}
		for _, fi := range dir {
			//fmt.Println(fi)
			go doIt(workerID, ch, done, fi, wg)
		}

	} else {
		fmt.Printf(fi.Name())
	}
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer wg.Done()
	for {
		select {
		case m := <-ch:
			fmt.Printf("[%v] m => %v\n", workerID, m)
		case <-done:
			fmt.Printf("[%v] is done\n", workerID)
			return
		}
	}
}
