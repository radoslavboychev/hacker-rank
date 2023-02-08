package main

import (
	"log"
	"strings"
)

func main() {

	res := minimumNumber(11, "#HackerRank")
	log.Println(res)
}

func minimumNumber(n int32, password string) int32 {
	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"
	var nu, l, u, s bool
	var counter int32
	for _, char := range password {
		if strings.Contains(numbers, string(char)) && !nu {
			nu = true
			counter++
		}
		if strings.Contains(lower_case, string(char)) && !l {
			l = true
			counter++
		}
		if strings.Contains(upper_case, string(char)) && !u {
			u = true
			counter++
		}
		if strings.Contains(special_characters, string(char)) && !s {
			s = true
			counter++
		}
	}
	add := 4 - counter
	if len(password) >= 6 {
		return add
	}
	if 6-len(password) > int(add) {
		return int32(6 - len(password))
	} else {
		return add
	}

}
