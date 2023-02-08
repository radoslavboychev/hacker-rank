package main

import "fmt"

func main() {
	res := arraySum([]int32{1, 2, 3, 4, 5})
	fmt.Println(res)
}

func arraySum(input []int32) int32 {
	var sum int32
	for _, v := range input {
		sum += v
	}
	return sum
}
