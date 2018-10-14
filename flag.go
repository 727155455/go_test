package main

import (
	"flag"
	"fmt"
)

var a *string = flag.String("aaa","bbb","ccc")
func main() {
	fmt.Println(*a)
	fmt.Printf("%+v ,,,,, %T", a, a)
	flag.Parse()
	fmt.Println(*a)
}