package main

import (
	"fmt"
	"unicode"
)

func main () {
	s := "𝕞𝕒𝕥𝕝𝕠𝕠𝕓"
	fmt.Println(s)
	for _, r := range s {
		fmt.Printf("unicode.IsLetter(%q): %v\n", r, unicode.IsLetter(r))
	}
}
