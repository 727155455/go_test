package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool, 10)

	for i:=0; i<10;i++ {
		go  mychan(i, i+1, channel)
	}

	for i:=0;i<10;i++ {
		<- channel
	}
	close(channel)
}

func mychan(a, b int, channel chan bool) {
	c := a + b
	fmt.Println(a, "+", b, "=", c)
	time.Sleep(time.Second * 2)
	channel <- true
}