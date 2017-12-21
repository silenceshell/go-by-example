package main

import (
	"fmt"
	"io"
	//"os"
	"github.com/labstack/gommon/bytes"
)

type xxx_i interface {}
type xxx_s struct{}

func main() {
	var x_i xxx_i
	if x_i == nil {
		fmt.Println("x_i is nil")
		// output is "x_i is nil"
	}

	var x_string string
	fmt.Println("begin" + x_string + "end")

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

	var w io.Writer
	check_nil(w)	//w is nil

	//var x os.Stdout
	var x *bytes.Bytes
	check_nil(x)	// x is not nil, because bytes.Bytes is a struct, x will contain object's type.
}

func check_nil(i interface{}) {
	fmt.Println("is is nil?", i==nil)
}
