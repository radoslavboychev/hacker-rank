package main

import (
	"fmt"
)

func main() {

	// var delta int
	// var input string
	// fmt.Scanln("Delta %d\n", &input)
	// fmt.Scanln("Input string %d\n", &delta)

	res := caesarCipher("There's a starman waiting in the sky", 3)

	fmt.Println(res)
}

// caesarCipher shifts each letter by a number of letters
func caesarCipher(s string, k int32) string {

	// create a new slice of runes to hold the unicode characters
	var alphabetSlice []rune

	// iterate through string, add each letter to the slice
	for _, letter := range s {
		alphabetSlice = append(alphabetSlice, letter)
	}

	// create a new slice for the "encrypted" string
	var res []rune
	// change the value of each unicode character
	for _, r := range alphabetSlice {

		// // change the value of the unicode character
		// r = r + k

		// if its larger than the 'Z' character, set it to 'A'
		// this loops the alphabet
		r = cipher(r, k)

		// append the shifted unicode rune to the string
		res = append(res, r)
	}

	// return the final string
	return string(res)
}

//
func cipher(r rune, delta int32) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', int(delta))
	}
	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', int(delta))
	}
	return r
}

//
func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}
