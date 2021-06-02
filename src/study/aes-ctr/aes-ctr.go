package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	fmt.Println("--------------")
	fmt.Println(GenSimplePushServerAuthToken("10009", "10000009"))
	fmt.Println("--------------")
	type GetSwitchStatus struct {
		DeviceSn  string `json:"deviceSn"`  //音箱sn
		TyPhone   string `json:"tyPhone"`   //天翼账号
		Timestamp int64  `json:"timestamp"` //请求时间戳，单位毫秒
	}



	type SyncSwitchStatus1 struct {
		DeviceSn     string `json:"deviceSn"`     //音箱sn
		TyPhone      string `json:"tyPhone"`      //天翼账号
		Timestamp    int64  `json:"timestamp"`    //请求时间戳，单位毫秒
		SwitchType   int    `json:"switchType"`   //1 回家看看，2 一呼即通，3 智能抓拍
		SwitchStatus int    `json:"switchStatus"` //1 开, 2 关
		External int `json:external` //
	}

	type DoorbellStateChangeReqBody struct {
		DeviceSnList string `json:"deviceSnList"`
		TpAccount string `json:"tpAccount"`
		DoorbellId string `json:"doorbellId"`
		Src string `json:"src"`
		State int `json:"state"`
		ReqTime int64 `json:"reqTime"`
		Kkkkk int `json:"kkkkk"`
	}


	type SyncSwitchStatus struct {
		DeviceSn     string `json:"deviceSn"`     //音箱sn
		TyPhone      string `json:"tyPhone"`      //天翼账号
		Timestamp    int64  `json:"timestamp"`    //请求时间戳，单位毫秒
		SwitchType   int    `json:"switchType"`   //1 回家看看，2 一呼即通，3 智能抓拍
		SwitchStatus int    `json:"switchStatus"` //1 开, 2 关
	}



	getSwitchStatus := GetSwitchStatus{
		DeviceSn: "C4C00400FF040310008K20149DB5A2BE20",
		TyPhone: "13369207073",
		Timestamp: time.Now().UnixNano() / 1e6,
	}
	getSwitchStatusBytes, _ := json.Marshal(getSwitchStatus)
	getSwitchStatusStr, _ := AesCTREncrypt(string(getSwitchStatusBytes), AES_CTR_KEY)
	println(getSwitchStatusStr)

	getSwitchStatusDecryptStr, _ := AesCTRDecrypt(getSwitchStatusStr, AES_CTR_KEY)
	println(getSwitchStatusDecryptStr)


	//{"tpAccount":"13305770106","src":"ctcc","deviceSnList":"C4C00400FF040361018T2115A1AA5A758A","state":2,"reqTime":1622618847331,"doorbellId":"3QHLB142154WCFU"}

	println("----------11111111111111111111--------")
	stateChange := DoorbellStateChangeReqBody{
		DeviceSnList:"11111111",
		TpAccount: "2222",
		DoorbellId: "333",
		Src : "444",
		State: 4,
		ReqTime: 5,
		Kkkkk: 7,
	}
	syncSwitchStatus := SyncSwitchStatus1{
		DeviceSn: "C4C00400FF040310008K20149DB5A2BE20",
		//DeviceSn: "CTCC-6H19243D9F169F2E",
		TyPhone: "13369207073",
		Timestamp: time.Now().UnixNano() / 1e6,
		SwitchStatus: 2,
		SwitchType: 2,
		External: 1,
	}
	syncSwitchStatusBytes, _ := json.Marshal(syncSwitchStatus)
	syncSwitchStatusStr, _ := AesCTREncrypt(string(syncSwitchStatusBytes), AES_CTR_KEY)
	println(syncSwitchStatusStr)

	stateChangeBytes, _ := json.Marshal(stateChange)
	stateChangeStr, _ := AesCTREncrypt(string(stateChangeBytes), AES_CTR_KEY)
	fmt.Println("99999999999")
	fmt.Println(stateChangeStr)
	fmt.Println("99999999999")

	syncSwitchStatusDecryptStr, _ := AesCTRDecrypt(syncSwitchStatusStr, AES_CTR_KEY)

	fmt.Println(syncSwitchStatusDecryptStr)
	aaaaaa := SyncSwitchStatus{}
	json.Unmarshal([]byte(syncSwitchStatusDecryptStr), &aaaaaa)
	fmt.Println(aaaaaa.DeviceSn)


	println("----------11111111111111111111--------")

	testStr, _ := AesCTRDecrypt("kfotF/SLa6YraSBQnemnrbVJg/hLg+UhoAkUpP9SyepQ0hyNzjOuweUKQwOMFEHz0jPqyiFcTgCPR1JfP0Wti1mra/m3+O/K54895egtJsZxtBsMS2MDvZuUU3ut1/kcH+Hp", AES_CTR_KEY)
	println(testStr)

	testStr1, _ := AesCTRDecrypt("kfotF/SLa6YraSBQnemnrbVJg/hLg+UhoAkUpP9SyepQ0hyNzjOuweUKQwOMFEHz0jPqyiFcTgCPR1JfP0Wti1mra/m3+O/K54895egtJsZxtBsMS2MDvZuUU3Ov0f0THe+4gq0K15lpRCa7j2z8IHVBFPZzdFPdvsufGJp3vtS/kN8=", AES_CTR_KEY)
	println(testStr1)
	println("---fdafda---------")

	testStr2, _ := AesCTRDecrypt("kfogAcSNeqERY2YP0Yip3Klb2KoIoNFn8Ut04a4WjKE52h2Q2G6Y7KACHxK9RwWkjTOkgl0WUgOLFxw+bwbqyBL8efS2sg==", AES_CTR_KEY)
	println(testStr2)

	println("------------")


	value, _ := strconv.Atoi("3244206695")
	println(value)

	testFunc(value)

	fmt.Println("各int类型的取值范围为：")
	fmt.Println("int8:", math.MinInt8, "~", math.MaxInt8)
	fmt.Println("int16:", math.MinInt16, "~", math.MaxInt16)
	fmt.Println("int32:", math.MinInt32, "~", math.MaxInt32)
	fmt.Println("int64:", math.MinInt64, "~", math.MaxInt64)
	fmt.Println()

}

func testFunc(aaa int) {
	fmt.Println(aaa)
}



const (
	AES_CTR_IV  = "7EtZy2zMlyBB6MNv"
	AES_CTR_KEY = "r6Je8tDyqBD7zqcv"
)

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

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

