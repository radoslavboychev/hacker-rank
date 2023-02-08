package main

import "fmt"

func main() {
	rest := caesarCipher("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 2)
	fmt.Println(rest)
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
	var i int32 = 0
	// change the value of each unicode character
	for _, r := range alphabetSlice {

		// change the value of the unicode character
		r = r + k
		i++
		// if its larger than the 'Z' character, set it to 'A'
		// this loops the alphabet

		if r < 65 {
			r = 65 + i
		}

		if r >= 91 {
			r = 65 + i
		}

		// if its larger than the 'z' character, set it to the 'a'
		// this loops the alphabet

		// append the shifted unicode rune to the string
		res = append(res, r)
	}

	// return the final string
	return string(res)
}
