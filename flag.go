package main

import (
	"flag"
	"fmt"
)

var a *string = flag.String("aaa", "bbb", "ccc")

func main() {
	fmt.Println(*a, a)
	fmt.Printf("%+v ,,,,, %v\n", a, a)
	flag.Parse()
	fmt.Println("----", *a, a)
}
