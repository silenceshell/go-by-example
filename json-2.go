package main

import (
    "encoding/json"
    "log"
    "os"
)

func main() {
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    for {
//        var v map[string]interface{}
	var v map[string]*json.RawMessage

        if err := dec.Decode(&v); err != nil {
//	if err := json.Unmarshal(data, &objmap)
            log.Println(err)
            return
        }
	log.Println("==========")
        for k,vv := range v {
		log.Println(k, vv)
            //if k != "Name" {
            //    delete(v, k)
            //}
        }
	log.Println("==========")
//	log.Println(v["spec"].(map[string]interface{})["kubernetesApiEndpoints"].(map[string]interface{})["serverEndpoints"].(map[string]interface{})["serverAddress"])
        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
    }
}
