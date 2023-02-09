package main

import (
	"fmt"
)

func main() {

	str := "hereiamstackerrank"
	res := hackerRankInString(str)
	fmt.Println(res)

}

func hackerRankInString(s string) string {
	target := "hackerrank"
	var currentIndex int = 0
	for i := 0; i < len(s); i++ {
		if s[i] == target[currentIndex] {
			currentIndex++
			if currentIndex == len(target) {
				return "YES"
			}
		}
	}
	return "NO"
}
