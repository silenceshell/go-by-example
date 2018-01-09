package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type org struct {
	name string
	address string
}

func main() {
	// download oui.csv from https://standards.ieee.org/develop/regauth/oui/oui.csv
	f, err:= os.Open("oui.csv")
	if err != nil {
		panic("file open failed")
	}
	defer f.Close()

	r := csv.NewReader(f)
	var macVendorHashMap map[string]org = make(map[string]org);
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var o org = org{name:record[2], address:record[3]}
		macVendorHashMap[record[1]] = o
	}

	mac := "1C:c1:De:99:88:77"
	newMac := strings.Replace(strings.ToUpper(mac), ":", "", -1)
	fmt.Println(newMac[:6])
	begin := time.Now()
	if v, ok := macVendorHashMap[newMac[:6]]; ok {
		fmt.Println(v.name, v.address)
	}
	fmt.Println(time.Now().Sub(begin))

}

