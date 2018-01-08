package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	defer fmt.Println("defer main")

	var user = os.Getenv("USER_")
	var ch chan int = make(chan int, 1)

	go func() {
		defer func() {
			fmt.Println("defer out loop")
		}()

		defer func() {
			fmt.Println("defer out loop 2")
		}()
		go func() {
			defer func() {
				fmt.Println("defer caller")

				if err := recover(); err != nil {
					fmt.Println("recover success.")
				}
			}()
			func() {
				defer func() {
					fmt.Println("defer here")

					ch <- 1
				}()

				defer func() {
					fmt.Println("defer here 2")
				}()
				if user == "" {
					//os.Exit(-1)
					panic("should set user env.")
				}
			}()
		}()
		time.Sleep(time.Second)
	}()

	result := <-ch
	//time.Sleep(1 * time.Second)
	fmt.Printf("get result %d\r\n", result)

}
