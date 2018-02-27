package main

import "fmt"

func main() {
	var x int = 3
	fmt.Println(x)

	x, y := 1, 3	//redeclare x
	fmt.Println(x, y)	//x is 1 now

	if true {
		x, z := 2, 4	//x is shadowed
		fmt.Println(x, z)	//x in this block is 2 now
	}

	fmt.Println(x)	//the original x is still 1
}
