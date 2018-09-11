package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i:=0;i<10;i++ {
		a[i] = i
	}


	fmt.Println(a)

	fmt.Println(a[10:])

	a = append(a[:9], a[10:]...)

	fmt.Println(a)

}
