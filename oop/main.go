package main

import (
	"fmt"
	"test/oop/employee"
)

func init() {
	fmt.Printf("init----\n")
}

func main() {
	//var a employee.Employee = new(employee.Employee)
	/*a := employee.Employee{
		Fn: "hehe",
		Ln: "xixi",
		Age: 18,
	}
	fmt.Printf("%+v ----", a)
	a.Abc()
	fmt.Printf("%+v ----", a)
	a.Abc()
	a.Age = 10
	fmt.Printf("%+v ----", a)*/
	a := employee.New("a", "b", 18, 10)
	fmt.Printf("%+v ----", a)
	a.Abc()
	fmt.Printf("%+v ----", a)
	a.SetAge(100)
	a.Abc()
	fmt.Printf("%+v ----", a)
}