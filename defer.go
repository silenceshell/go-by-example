package main

import "fmt"

func main(){
    a :="1"
    defer pring(a)
    a ="2"
}
func pring(a string){
    fmt.Println(a)
}
