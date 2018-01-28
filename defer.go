package main

import "fmt"

func main(){
    a :="1"
    defer print(a)
    a ="2"

    ret := call_test()
    fmt.Println(ret)
}
func print(a string){
    fmt.Println(a)
}

func call_test() (ret int) {
    defer func(){
        ret = 3
    }()

    //ret = 1
    return 1
}