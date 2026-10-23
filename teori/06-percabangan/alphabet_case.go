package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string

	fmt.Scanf("%s", &input)

	if len(input) > 0 {

		karakter := rune(input[0])

		if unicode.IsLetter(karakter) {
			fmt.Println("Alphabet")
		} else {
			fmt.Println("Bukan Alphabet")
		}
	}
}
