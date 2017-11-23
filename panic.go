package main

import (
	"os"
	"fmt"
)

func main() {

	var user = os.Getenv("USER_")
	var ch chan int = make(chan int, 1)

	go func() {
		defer func() {
			fmt.Println("defer here")
			if err := recover(); err != nil {
				fmt.Println("recover success.")
			}

			ch <- 1
		}()

		if user == "" {
			panic("should set user env.")
		}
	}()

	result := <- ch
	fmt.Printf("get result %d\r\n", result)

}
