package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"strconv"
)

func main() {
	spew.Dump(strconv.FormatInt(3600*2*1000, 10))
	spew.Dump(string(3600*2*1000))


	fmt.Println(MD5("fdafd地对大法师地飞弹fdsa"))
}



func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

