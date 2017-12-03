package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 33333

	// wrong!
	var xs = string(x)
	fmt.Println(xs)

	xs = strconv.Itoa(x)
	fmt.Println(xs)
}
