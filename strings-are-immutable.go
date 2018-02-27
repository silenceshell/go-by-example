package main

import "fmt"

func main() {
	var x string = "hello world"
	fmt.Println(x)

	//x[0] = "h" //compile error

	var xb []byte = []byte(x)
	xb[0] = 'H'
	fmt.Println(x)
	fmt.Println(string(xb), xb)

	fmt.Printf("%T", x[0])
}
