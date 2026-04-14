package main

import (
	"fmt"
	"time"
)

func sayHello(ch chan string) {
	fmt.Println(<-ch)
}

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)
	// go sayHello(channel)
	go sayHello2(channel1, channel2)
	channel1 <- "Goroutine 1"
	fmt.Println("Main")
	// time.Sleep(2 * time.Second)
	channel2 <- "Goroutine 2"
	// time.Sleep(2 * time.Second)
	time.Sleep(1 * time.Second)
}

func sayHello2(ch1 chan string, ch2 chan string) {
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
