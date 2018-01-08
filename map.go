package main

import (
	"fmt"
	//"sort"
	"sort"
)

func main() {
	xx := make(map[string]string)

	xx["111"] = "xxx"
	xx["222"] = "zzz"

	for v := range xx {
		fmt.Println(v)
	}

	zz := make([]string, 0, 3)
	zz = append(zz, "v1/Secret")
	zz = append(zz, "v1beta1/Deployment")
	zz = append(zz, "v1/Pod(related)")
	zz = append(zz, "v1/Service")

	sort.Strings(zz)

	for _, v := range zz {
		fmt.Println(v)
	}
}
