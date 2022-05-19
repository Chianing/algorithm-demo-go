package test

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
)

const (
	jobCnt     = 100
	routineCnt = 5
)

// 控制线程数量
func routinePoolExecute() {
	ch := make(chan bool, routineCnt)
	retCh := make(chan int, jobCnt)
	// wg := sync.WaitGroup{}
	// wg.Add(jobCnt)

	fmt.Printf("current goroutine cnt: %d\n", runtime.NumGoroutine())
	fmt.Println()

	for i := 0; i < jobCnt; i++ {
		subRet := i
		ch <- true

		go func() {
			defer func() {
				<-ch
			}()
			// defer wg.Done()
			fmt.Printf("goroutine cnt: %d, go func: %d\n", runtime.NumGoroutine(), subRet)
			retCh <- subRet
		}()
	}

	// wg.Wait()

	ret := make([]int, 0)
	for i := 0; i < jobCnt; i++ {
		ret = append(ret, <-retCh)
	}

	sort.Ints(ret)
	fmt.Printf("got ret: %v", ret)

}

// 两个协程交替打印AB
func printInOrder() {
	flgA := make(chan bool, 1)
	flgB := make(chan bool, 1)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			<-flgB
			fmt.Print("A")
			flgA <- true
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-flgA
			fmt.Println("B")
			flgB <- true
		}
		wg.Done()
	}()

	flgB <- true

	wg.Wait()

}
