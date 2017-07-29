package main

import "fmt"

type xxx_i interface {}
type xxx_s struct{}

func main() {
	var x_i xxx_i
	if x_i == nil {
		fmt.Println("x_i is nil")
		// output is "x_i is nil"
	}

	var x_s xxx_s

	// uncomment below codes, you will get an compile error: cannot convert nil to type xxx_s
	// because var_s is a structure instance, you can compare its pointer to nil.
	//if x_s == nil {
	//	fmt.Println("x_s is nil")
	//}

	// compare like this.
	if &x_s != nil {
		fmt.Println("x_s is not nil")
	}
}


