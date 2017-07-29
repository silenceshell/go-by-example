package main

import (
	"fmt"
	"encoding/json"
	//"github.com/gin-gonic/gin/json"
	"reflect"
	"strings"
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
func json_example() {
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
		fmt.Printf("user b name %s, passwd %s\r\n", user_b_tag.Name, user_b_tag.Passwd)
		//should be user b name lily, passwd 123456
	}
}

type custom struct {
	Name 	string 	`custom:"custom_name"`
	Passwd 	string	`custom:"custom_password"`
}

func (obj * custom) Fmt() string {
	var ret string
	var ret_array []string
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	te := t.Elem()
	ve := v.Elem()
	for i:=0; i< te.NumField(); i++ {
		ret_array = append(ret_array, fmt.Sprintf(`"%s": "%v" `, te.Field(i).Tag.Get("custom"), ve.Field(i).Interface()))
	}
	ret = strings.Join(ret_array, ",")
	return `{` + ret + `}`
}

func custom_example() {
	user_c := &custom{"lucy", "666666"}
	fmt.Println(user_c.Fmt())
}

func main() {
	json_example()
	custom_example()

}
