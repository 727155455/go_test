package main

import (
	"fmt"
	"sync"
)

func main() {
	var go_sync sync.WaitGroup
	aa := 0
	for i:=0; i<10;i++ {
		go_sync.Add(1)
		go  mydata(i, &aa, &go_sync)
	}
	go_sync.Wait()
	fmt.Println("end aa : ", aa)
}

func mydata(a int, aa *int, g *sync.WaitGroup) {
	//c := a + b
	//fmt.Println(a, "+", b, "=", c)
	*aa += a
	fmt.Println("routine i :" ,a , " aa : ", *aa)
	defer g.Done()
}