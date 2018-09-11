package main

import "fmt"

func main() {
	var x []string = make([]string, 0, 10)

	x = append(x, "abc")
	x = append(x, "b")
	fmt.Println(len(x))
	fmt.Println(x[0], x[1])

	var xx []byte = make([]byte, 10)
	xx[0]='a'
	xx[1]='"'
	xx[2] = 'a'
	var sxx = string(xx)
	fmt.Println(xx)
	fmt.Println(sxx)

	x2 := "Bearer eyJhbGciOiJS"
	fmt.Println(x2)
	fmt.Println(x2[7:])
}
