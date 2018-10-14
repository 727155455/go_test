package main

import "fmt"

func main(){
	var arr = make([]int, 4)
	for i:=0; i<10; i++ {
		go func(arr[]int, i int){
			/*defer func() {
				if err := recover(); err != nil {
					fmt.Println("panic : " , err)
				}
			}()*/
			fmt.Println("i : " , i)
			createPanic(arr, i)
		}(arr, i)

	}
	fmt.Println("end")
}

func createPanic(arr []int, i int){
	fmt.Println("第", i, "次--->", arr[i])
}
