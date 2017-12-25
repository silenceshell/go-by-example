package main

import "fmt"

type Request struct {
	args []int
	f func([]int) int
	ch chan int
}

func sum(nums []int) int {
	tmp := 0
	for _, v := range nums {
		tmp += v
	}
	return tmp
}

var maxOutStanding int = 10

func handle(requests chan *Request) {
	for req := range requests {
		req.ch <- req.f(req.args)
	}
}

func Serve(requests chan *Request) {
	for i:=0; i<maxOutStanding; i++ {
		go handle(requests)
	}
}

func main() {
	var requests chan *Request = make(chan *Request, 3)

	Serve(requests)

	req := &Request{[]int{1,3,4}, sum, make(chan int)}
	requests <- req
	fmt.Printf("result %d\n", <- req.ch)

}
