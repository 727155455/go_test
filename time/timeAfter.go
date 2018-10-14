package main

import (
	"fmt"
	"time"
)

func main(){
	time.AfterFunc(2 * time.Second, func(){
		fmt.Println("running2    ", time.Now())
	})
	chanTime := time.After(3 * time.Second)
	fmt.Println("start    ", time.Now() )
	_ = <- chanTime
	fmt.Println("running1    ", time.Now())

	fmt.Println("end    ", time.Now())

	fmt.Println("start2    ", time.Now() )
	time.AfterFunc(2 * time.Second, func(){
		fmt.Println("running2    ", time.Now())
	})
	chanTime3 := time.After(4 * time.Second)
	fmt.Println("start3    ", time.Now() )
	_ = <- chanTime3
	fmt.Println("end2    ", time.Now())


}