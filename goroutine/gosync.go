package main

import (
	"fmt"
	"sync"
)

func main() {
	var go_sync *sync.WaitGroup
	for i:=0; i<10;i++ {
		go_sync.Add(1)
		go  mysync(i, i+1, go_sync)
	}
	go_sync.Wait()
}

func mysync(a, b int, g *sync.WaitGroup) {
	c := a + b
	fmt.Println(a, "+", b, "=", c)
	defer g.Done()
}