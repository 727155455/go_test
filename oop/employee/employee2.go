package employee

import "fmt"

type employee2 struct {
	fn string
	ln string
	age int
	e int
}

func New(fn, ln string, age, e int) *employee2 {
	ee := &employee2{fn, ln, age, e}
	return ee
}

func init() {
	fmt.Printf("init2----\n")
}

func (e employee2) Abc() {
	e.e++
	fmt.Println(e.age,e.fn,e.ln,e.e)
	e.age++
}

func (e employee2) SetAge(newAge int) {
	e.age = newAge
}