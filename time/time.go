package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	debug.SetTraceback("system")
	/*if len(os.Args) == 1 {
		panic("before timers")
	}*/
	for i := 0; i < 10; i++ {
		time.AfterFunc(time.Duration(5*time.Second), func() {
			fmt.Println("Hello!")
		})
	}
	panic("after timers")
}
