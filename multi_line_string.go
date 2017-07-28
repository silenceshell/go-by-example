package main

import "fmt"

func main() {
	str1 := `{"user_name":"lily","user_passwd":"123456"}`
	fmt.Println(str1)

	fmt.Println("------------")

	str2 :=
`line1
line2
line3`
	fmt.Println(str2)

	fmt.Println("------------")

	str3 := "line1\nline2\nline3\n"
	fmt.Println(str3)
}
