package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data := "♥"
	fmt.Println(len(data))	//should be 3
	fmt.Println(utf8.RuneCountInString(data)) //should be 1
	fmt.Println(utf8.RuneCountInString("é")) //should be 1

	data = "é"
    fmt.Println(len(data))                    //prints: 3
    fmt.Println(utf8.RuneCountInString(data)) //prints: 2

	fmt.Println(utf8.ValidString("ABC")) //should be true
	fmt.Println(utf8.ValidString("X\xfeY")) //should be false
}
