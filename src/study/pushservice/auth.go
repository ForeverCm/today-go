package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)


const (
	PUSH_SERVICE_AUTH_NAME   = "10017"
	PUSH_SERVICE_AUTH_SECRET = "10000017"
	PATH_PUSH_MESSAGE        = "/message/v2/pushmessage"
)



func main() {
	fmt.Println(GenAuthToken(PUSH_SERVICE_AUTH_NAME, PUSH_SERVICE_AUTH_SECRET))
}

func GenAuthToken(name string, secret string) string {
	timeStr := strconv.FormatInt(time.Now().Unix(), 10)
	md5Str := GetMD5Hash(name + "-" + secret + "-" + timeStr)
	token := name + "-" + timeStr + "-" + md5Str
	return token
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
