package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	text := "hhhhhh"
	hasher := md5.New()
	hasher.Write([]byte(text))

	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))

}