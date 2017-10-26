package main

import (
	"fmt"
	"reflect"
)

func main() {

    v := "hello world"
    fmt.Println(reflect.TypeOf(v).String())

    fmt.Printf("%T", v)
	fmt.Println(typeof(v))
}

func typeof(v interface{}) string {
    switch t := v.(type) {
    case int:
        return "int"
    case float64:
        return "float64"
	case string:
		return "string"
    //... etc
    default:
        _ = t
        return "unknown"
    }
}