package main

import "fmt"

func main() {
	var x []string = make([]string, 0, 10)

	x = append(x, "abc")
	x = append(x, "b")
	fmt.Println(len(x))
	fmt.Println(x[0], x[1])

}