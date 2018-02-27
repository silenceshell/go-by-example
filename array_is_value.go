package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x [3]int = [3]int {1, 2, 3}
	fmt.Println(reflect.TypeOf(x))

	fmt.Println(x)

	// pass slice x as a value to func
	func (xx [3]int) {
		xx[0] = 11
		fmt.Println(xx)
	}(x)

	fmt.Println(x)

	// pass pointer of x to func
	func (xx *[3]int) {
		xx[0] = 11
	}(&x)

	fmt.Println(x)

	// slice is not array.
	var x2 []int = []int {1, 2, 3}
	//var x2 []int = make([]int, 3)
	//x2[0], x2[1], x2[2] = 1, 2, 3
	fmt.Println(reflect.TypeOf(x2))

	func (xx []int) {
		xx[0] = 11
	}(x2)
	fmt.Println(x2)

}
