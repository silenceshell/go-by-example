package main

import (
    flag "github.com/spf13/pflag"
    "fmt"
)
// go run pflag.go --name 998 --passwd=999
func main() {
	var ip *int = flag.Int("name", 1234, "help message for flagname")
	var flagvar int
	flag.IntVar(&flagvar, "passwd", 1234, "help message for flagname")

	flag.Parse()

	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)
}
