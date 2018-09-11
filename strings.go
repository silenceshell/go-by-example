package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "1.1.1.1:2222"
	arr := strings.Split(a, ",")
	fmt.Println(len(arr))

}
