package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func main() {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)
	fmt.Println(dat["num"])

	macAddress := "08:74:02:00:00:00"
	resp, err := http.Get(fmt.Sprintf("https://macvendors.co/api/%s/", macAddress))
	defer resp.Body.Close()
	if err!= nil {
		panic(err)
	}

	//result := []byte(`{"result":{"company":"Apple, Inc.","mac_prefix":"08:74:02","address":"1 Infinite Loop,Cupertino  CA  95014,US","start_hex":"087402000000","end_hex":"087402FFFFFF","country":"US","type":"MA-L"}}`)
	result, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
		panic(err)
	}
	var rat map[string]map[string]interface{}
	if err := json.Unmarshal(result, &rat); err != nil {
        panic(err)
    }
    fmt.Println(rat)
	fmt.Println(rat["result"]["company"])
	fmt.Println(rat["result"]["address"])



}
