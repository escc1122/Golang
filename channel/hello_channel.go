package main

import (
	"fmt"
	"time"
)

// 練習基本的channel
func hello() {
	var stringChen = make(chan string, 3)
	stringChen <- "hello"
	stringChen <- "channel"
	fmt.Println(<-stringChen)
	fmt.Println(<-stringChen)
}

// channel & Goroutine
func hello2() {
	var stringChen = make(chan string, 3)
	go func() {
		time.Sleep(1000 * time.Millisecond)
		stringChen <- "hello"
	}()

	go func() {
		time.Sleep(5000 * time.Millisecond)
		stringChen <- "world"
	}()

	fmt.Println(<-stringChen)
	fmt.Println(<-stringChen)
}

// channel & range
func withRange() {
	c := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // 關閉 Channel
	}()
	for i := range c { // 在 close 後跳出迴圈
		fmt.Println(i)
	}
}

// channel & select
func withSelect() {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("calculate goroutine starts calculating")
			time.Sleep(time.Second) // Heavy calculation
			fmt.Println("calculate goroutine ends calculating")

			ch <- "FINISH"
			time.Sleep(time.Second)
			fmt.Println("calculate goroutine finished")
		}
	}()

	for {
		select {
		case <-ch: // Channel 中有資料執行此區域
			fmt.Println("main goroutine finished")
			return
		default: // Channel 阻塞的話執行此區域
			fmt.Println("WAITING...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// hello()
	//hello2()
	// withRange()
	withSelect()
}
