package main

import "fmt"

func main() {
	res := arraySum([]int64{1, 2, 3, 4, 5})
	fmt.Println(res)
}

func arraySum(input []int64) int64 {
	var sum int64
	for _, v := range input {
		sum += v
	}
	return int64(sum)
}
