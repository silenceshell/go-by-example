package main

import (
	"fmt"
)

func main() {

	xs := make([]string, 0)
        xs = append(xs, "hhh")
	for index, x := range xs {
		fmt.Println(index)
		fmt.Println(x)
	}
	xs = nil
	for index, x := range xs {
		fmt.Println(index)
		fmt.Println(x)
	}

}

