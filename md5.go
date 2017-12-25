package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	text := "github.com/cw1997/TinyUrl/blob/master/src/url/url.go"
	hasher := md5.New()
	hasher.Write([]byte(text))
/*
1)将长网址md5生成32位签名串,分为4段, 每段8个字节;
2)对这四段循环处理, 取8个字节, 将他看成16进制串与0x3fffffff(30位1)与操作, 即超过30位的忽略处理;
3)这30位分成6段, 每5位的数字作为字母表的索引取得特定字符, 依次进行获得6位字符串;
4)总的md5串可以获得4个6位串; 取里面的任意一个就可作为这个长url的短url地址;
*/
	var alphabet string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
	sum := hasher.Sum(nil)
	sumHex := hex.EncodeToString(hasher.Sum(nil))
	fmt.Println(sum)
	fmt.Println(sumHex)

	var x [][]string = make([][]string, 4, 4)

	for i:=0; i<4; i++ {
		x[i] = make([]string, 0, 6)
		tmp := sumHex[i*8:(i+1)*8]
		tmpInt, err := strconv.ParseInt(string(tmp), 16, 64)
		fmt.Printf("%x\r\n", tmpInt)
		if err != nil {
			fmt.Println(err.Error())
		}
		tmpInt = tmpInt & 0x3fffffff

		var cInt int64
		for j:=0; j<6; j++ {
			cInt = int64(tmpInt) & 0x1f
			tmpInt = tmpInt >> 5
			x[i] = append(x[i], string(alphabet[cInt]))
		}
	}

	fmt.Println(x)
}