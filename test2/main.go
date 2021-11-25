// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"fmt"
	"log"
	"time"

	"test2/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

func test1() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	urls := []string{"https://golang.org", "https://go.dev/solutions/paypal"}

	n++
	go func() { worklist <- urls }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func test2() {
	worklist := make(chan []string, 20)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	urls := []string{"https://golang.org"}

	n = 3
	go func() { worklist <- urls }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}

	for {
		select {
		case <-worklist: // Channel 中有資料執行此區域
			fmt.Println("main goroutine finished")
		default: // Channel 阻塞的話執行此區域
			fmt.Println("WAITING...")
			time.Sleep(500 * time.Millisecond)
		}
	}

}

//!+
func main() {
	test2()
}

//!-
