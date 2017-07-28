package main

import (
	"fmt"
	"encoding/json"
	//"github.com/gin-gonic/gin/json"
)

type user struct {
	Name   string
	Passwd string
}

type user_tag struct {
	Name   string		`json:"user_name"`
	Passwd string	`json:"user_passwd"`
}

// struct tag is useful for json

func main() {
	user_a := &user{"lily", "123456"}
	user_a_s, err := json.Marshal(user_a)
	if err == nil {
		fmt.Println(string(user_a_s))
		// should be {"Name":"lily","Passwd":"123456"}
	}

	user_a_tag := &user_tag{"lily", "123456"}
	user_a_tag_s, err := json.Marshal(user_a_tag)
	if err == nil {
		fmt.Println(string(user_a_tag_s))
		// should be {"user_name":"lily","user_passwd":"123456"}
	}

	user_b_s := `{"user_name":"lily","user_passwd":"123456"}`
	user_b_tag := &user_tag{}
	if json.Unmarshal([]byte(user_b_s), user_b_tag) == nil {
		fmt.Printf("user b name %s, passwd %s", user_b_tag.Name, user_b_tag.Passwd)
		//should be user b name lily, passwd 123456
	}

}
