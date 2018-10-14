package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", twoSum([]int{2,7,11,15}, 9))
	fmt.Printf("%#v\n", twoSum([]int{2,7,11,15}, 10))
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for k, v := range nums {
		if i, ok := hash[target - v]; ok {
			return []int{i ,k}
		}
		hash[v] = k
	}
	return []int{0 ,0}
}