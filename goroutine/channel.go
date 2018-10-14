package main

import (
	"fmt"
	"time"
)

func main(){
	a := make(chan int)
	a <- 2
	go getChannel(a)
	time.Sleep(time.Second * 3)
}

func getChannel(a chan int){
	b, err := <- a;
	fmt.Println("get channel value : ", b ," err : ", err)
	b, err = <- a;
	fmt.Println("get2 channel value : ", b ," err : ", err)
}
