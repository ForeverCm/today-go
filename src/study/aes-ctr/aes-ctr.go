package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"strconv"
	"strings"
	"time"
)

const (
	AES_CTR_IV  = "7EtZy2zMlyBB6MNv"
	AES_CTR_KEY = "3sOzeOzdjhzv6EzmTHrE5w7LGpYT83VQ"

	Secret_AES_CONTACT_CTR_IV  = "7EtZy2zMlyBB6MNv"
	Secret_AES_CONTACT_CTR_KEY = "3sOzeOzdjhzv6EzmTHrE5w7LGpYT83VQ"

)

func main() {
	type GuHuaNotifyReq struct {
		Type         int    `json:"type"`
		TyAccount    string `json:"tyAccount"`
		CtccDeviceSn string `json:"ctccDeviceSn"`
		Ctei         string `json:"ctei"`
		BindKey      string `json:"bindKey"`
		Telephone    string `json:"telephone"`
		Timestamp    int64  `jons:"timestamp"`
	}

	guHuaNotifyReq := GuHuaNotifyReq{
		Type : 1,
		TyAccount: "18829026981",
		CtccDeviceSn: "6H19243D9F169F2E",
		Ctei : "kdkfsafd",
		BindKey : "abababab",
		Telephone : "18277778888",
		Timestamp : time.Now().Unix(),
	}



	type CtccGuhuaContactResDataAddrBook struct {
		Id     int64  `json:"id"`
		Name   string `json:"name"`
		Mobile string `json:"mobile"`
	}

	type CtccGuhuaContactResData struct {
		TyAccount string                            `json:"tyAccount"`
		BookList  []CtccGuhuaContactResDataAddrBook `json:"bookList"`
	}

	ctccGuhuaContactResDataAddrBook := CtccGuhuaContactResDataAddrBook{
		Id: 1,
		Name: "chenmeng",
		Mobile: "18829026981",
	}
	ctccGuhuaContactResDataAddrBook2 := CtccGuhuaContactResDataAddrBook{
		Id: 2,
		Name: "xiaoming",
		Mobile: "18829047896",
	}

	ctccGuhuaContactResData := CtccGuhuaContactResData{
		TyAccount: "18829026981",
		BookList: []CtccGuhuaContactResDataAddrBook{
			ctccGuhuaContactResDataAddrBook,
			ctccGuhuaContactResDataAddrBook2,
		},
	}

	fmt.Println("-----------999999999999-----------")
	getSwitchStatusBytesBytes, _ := json.Marshal(ctccGuhuaContactResData)
	getSwitchStatusStr, _ := AesCTREncrypt(string(getSwitchStatusBytesBytes), AES_CTR_KEY)
	println(getSwitchStatusStr)

	type QueryCtccGuHuaContactPerson struct {
		ContactId   int64  `json:"contactId"`
		Nickname    string `json:"nickname"`
		PhoneNumber string `json:"phoneNumber"`
	}
	type GetCtccGuHuaContactRes struct {
		Persons []*QueryCtccGuHuaContactPerson `json:"persons"`
		Version int64                          `json:"version"`
		Type    string                         `json:"type"`
	}
	queryCtccGuHuaContactPerson := QueryCtccGuHuaContactPerson{
		ContactId: 1,
		Nickname: "chenmeng",
		PhoneNumber: "18829026981",
	}
	queryCtccGuHuaContactPerson1 := QueryCtccGuHuaContactPerson{
		ContactId: 2,
		Nickname: "xiaoming",
		PhoneNumber: "18829047896",
	}

	getCtccGuHuaContactRes := []*QueryCtccGuHuaContactPerson{
		&queryCtccGuHuaContactPerson,
		&queryCtccGuHuaContactPerson1,
	}

	fmt.Println("-----------kkkkkkkkkkkkkkkkkkkkkkkkkk-----------")
	getCtccGuHuaContactResByte, _ := json.Marshal(getCtccGuHuaContactRes)
	getCtccGuHuaContactResStr, _ := AesCTREncrypt(string(getCtccGuHuaContactResByte), AES_CTR_KEY)
	println(getCtccGuHuaContactResStr)















	fmt.Println("-----------111111------------")
	getSwitchStatusBytes, _ := json.Marshal(guHuaNotifyReq)
	getSwitchStatusStr, _ = AesCTREncrypt(string(getSwitchStatusBytes), AES_CTR_KEY)
	println(getSwitchStatusStr)

	fmt.Println("-----------2222222------------")

	getSwitchStatusDecryptStr, _ := AesCTRDecrypt("kfo9C8OBa6wNaXZIhYii3bZIhfBP85EkoRsLt60NlrlXiV/I2DCx+PIiFWPzE13zk3782h1RA1TIVFtOP0SmiVarbfml4/6GqsAsrr9qsiagoRMMDH5OrsTFSXCsz+9JQrT9zLtfhM87H0HzzTHqLHVeBqcoP0nfu/3JQ8zlQ3392oAPPxE=", AES_CTR_KEY)
	fmt.Println(getSwitchStatusDecryptStr)


	fmt.Println("0000000000000000000000000000")
	fmt.Println(getTableNum("6H19243D9F169F2E"))

	fmt.Println(makeGuhuaPlatformSignature("2021-07-13 16:53:48", "/rtc/third/contact/query"))


	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))


	fmt.Println("1111111111")
	fmt.Println(AesCTRDecryptWithIv("kfo5F/CRZ60LJTgxxIjwgesN1qsPjMczrggLt6ELmrl1gUHZ2DDI4LguHyysSBbz0jPu2x5aRCCfCAoYfFSkn1GhY/y+/+7e8pV4rvhkLpBztgVaEDFBxcmDUXi0waNNTr36wbMYnNcoVBujkGS3dCBPGqd0dUjQs9aeFIxnv9S/gJNKWl5FJy4qXpOHgpQTIYSx4h9dByrIH2wdAeoZ9nMs+s04nJYVgmQGL3J/IG0=", Secret_AES_CONTACT_CTR_KEY, Secret_AES_CONTACT_CTR_IV))

}

func testFunc(aaa int) {
	fmt.Println(aaa)
}





func aesCtrCrypt(text []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if blockSize != len(iv) {
		return nil, errors.New("iv len invalid")
	}
	// 指定分组模式
	blockMode := cipher.NewCTR(block, iv)
	// 执行加密、解密操作
	message := make([]byte, len(text))
	blockMode.XORKeyStream(message, text)
	// 返回明文或密文
	return message, nil
}

func AesCTREncrypt(plainText string, key string) (string, error) {
	cipherText, err := aesCtrCrypt([]byte(plainText), []byte(key), []byte(AES_CTR_IV))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func AesCTRDecrypt(cipherText string, key string) (string, error) {
	plainText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	plainText, err = aesCtrCrypt([]byte(plainText), []byte(key), []byte(AES_CTR_IV))
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func GenSimplePushServerAuthToken(name string, secret string) string {
	timeStr := strconv.FormatInt(time.Now().Unix(), 10)
	md5Str := GetMD5Hash(name + "-" + secret + "-" + timeStr)
	token := name + "-" + timeStr + "-" + md5Str
	return token
}

func AesCTRDecryptWithIv(cipherText string, key string, ivStr string) (string, error) {

	plainText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	plainText, err = aesCtrCrypt([]byte(plainText), []byte(key), []byte(ivStr))
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}


func getTableNum(sn string) uint32 {
	return crc32.ChecksumIEEE([]byte(sn)) % 4
}


func makeGuhuaPlatformSignature(reqDate, requestURI string) string {
	param := fmt.Sprintf("AppId=%s&ReqDate=%s&RequestURI=%s&Param=%s", "88010012", reqDate, requestURI, "")
	key := []byte("ZJl3AZ67spXiaKpVd6SVxxNnYdlZpYeG")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(param))
	return strings.ToTitle(hex.EncodeToString(mac.Sum(nil)))
}


