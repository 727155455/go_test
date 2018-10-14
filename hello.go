package main

import "fmt"

func main() {
	fmt.Println("12345")
	defer fmt.Println("12347")
	fmt.Println("1234")
	defer fmt.Println("12346")
}
