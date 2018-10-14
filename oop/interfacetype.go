package main

import "fmt"

func main() {
	var data map[string]interface{}
	fmt.Printf("111111%+v\n", data)
	data = make(map[string]interface{})
	data["aaa"] = []int{1,2,3}
	data["ccc"] = []int{1,2,3}
	fmt.Printf("222222%+v\n", data)
	var b interface{}
	i:=1
	b=data
	fmt.Printf("555555%+v\n%+v\n", b,data)
	data["ddd"] = []int{1,2,3}
	//b["ddd"] = []int{1,2,3}
	fmt.Printf("3333333%+v\n%+v\n", b,data)

	b=i
	fmt.Println("b:", b, "i:", i)
	i=4
	fmt.Println("b:", b, "i:", i)
}
