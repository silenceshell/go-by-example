package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1 chan struct{} = make(chan struct{}, 1)
	var ch2 chan struct{} = make(chan struct{}, 1)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("will close ch1 and ch2 now.")
		close(ch1)
		close(ch2)
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case _, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 is closed.")
				ch1 = nil
			}
		case _, ok:= <- ch2:
			if !ok {
				fmt.Println("ch2 is closed.")
				ch2 = nil
			}
		}
	}

	//var ch1Closed bool = false
	//var ch2Closed bool = false
	//for ch1Closed != true || ch2Closed != true {
	//	select {
	//	case _, ok := <-ch1:
	//		if !ok {
	//			fmt.Println("ch1 is closed.")
	//			ch1Closed = true
	//		}
	//	case _, ok := <-ch2:
	//		if !ok {
	//			fmt.Println("ch2 is closed.")
	//			ch2Closed = true
	//		}
	//	}
	//}

}
