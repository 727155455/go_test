package employee

import "fmt"

type Employee struct {
	Fn string
	Ln string
	Age int
	e int
}

func init() {
	fmt.Printf("init0----\n")
}

func (e *Employee) Abc() {
	e.e++
	fmt.Println(e.Age,e.Fn,e.Ln,e.e)
	e.Age++
}