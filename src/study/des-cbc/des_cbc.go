package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"runtime"
)

func main() {
	type TestDes struct {
		Cuei        string `json:"cuei"` //设备sn编号
		Uid         string `json:"uid"`  //本地ctei，没有则传空
		Timestamp   int    `json:"timestamp"`
		SwitchState int    `json:"switchState"`
	}

	aaa := TestDes{
		Cuei:        "184085900000022",
		Uid:         "testtesttest",
		Timestamp:   1610363331123,
		SwitchState: 1,
	}

	aaaStr, _ := json.Marshal(aaa)

	fmt.Println(string(aaaStr))
	bbbb, _ := LTTripleDesEncrypt(string(aaaStr))

	fmt.Println(bbbb)

	cccccStr, _ := LTTripleDesDecrypt(bbbb)

	m := make(map[string]interface{})
	json.Unmarshal([]byte(cccccStr), &m)

	fmt.Println(m)
}





var (
	ErrCipherKey=errors.New("The secret key is wrong and cannot be decrypted. Please check")
	ErrKeyLengthSixteen=errors.New("a sixteen-length secret key is required")
	ErrKeyLengtheEight=errors.New("a eight-length secret key is required")
	ErrKeyLengthTwentyFour=errors.New("a twenty-four-length secret key is required")
	ErrPaddingSize=errors.New("padding size error please check the secret key or iv")
	ErrIvAes=errors.New("a sixteen-length ivaes is required")
	ErrIvDes=errors.New("a eight-length ivdes key is required")
)

const (
	ivdes = "aproblem"
	//CUCCKey = "033f6b86b6d722d23e15f5b77913d451" // 联通3des key，需密钥移位
	CUCCKey = "d616648c09ee5cc3de0cfb3a1b92b99c"  //测试环境联通key
)

//生产环境CUCCKey
//{"cuei":"184085900000022","uid":"testtesttest","timestamp":1610363331123,"switchState":1}
//zNq7vsJf8pAGxM6fsTQvoKi+MqbKOkbUrjaFieQZ3TNcBv3TSxDSsP8hXu6rY7PRtr9S6VQQTZ8F0d0xdvVYQbCYuTZyr+sIyqAva/P06pCg54ijnD07QHW7cbeVJ1pV
//map[cuei:184085900000022 switchState:1 timestamp:1.610363331123e+12 uid:testtesttest]

//测试环境CUCCKey
//{"cuei":"184085900000022","uid":"testtesttest","timestamp":1610363331123,"switchState":1}
//teIla+aR4+BMlYANhihY4UGLw5dhn8xXH0K1iqiCnuqKHqUmqb+iUK1twM+68SnxtJaehKtq9Zji4vNDCllkwR5QRkOq0vkM0vgBZ3JUL8uSMYAOg6CfxbpIFnRTpWqZ
//map[cuei:184085900000022 switchState:1 timestamp:1.610363331123e+12 uid:testtesttest]



// 联通加密方式，3des
func LTTripleDesEncrypt(plainStr string) (string, error) {
	plainText := []byte(plainStr)

	var iv int = 0
	ivBytes := make([]byte, 8)
	for i:= 0; i < 8 ; i++ {
		ivBytes[i] = byte(iv)
	}

	keyBytes, _ := GetKeyBytes([]byte(CUCCKey))

	// 传入明文和自己定义的密钥，密钥为24字节 可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	cryptText, err := TripleDesEncrypt(plainText, keyBytes, ivBytes...)
	if err != nil {
		return "", err
	}

	base64CryptStr := base64.StdEncoding.EncodeToString(cryptText)
	return base64CryptStr, nil
}

// 联通解密方式，3des
func LTTripleDesDecrypt(cryptStr string) (string, error){
	cryptText, err := base64.StdEncoding.DecodeString(cryptStr)
	if err != nil {
		return "", err
	}

	var iv int = 0
	ivBytes := make([]byte, 8)
	for i:= 0; i < 8 ; i++ {
		ivBytes[i] = byte(iv)
	}

	keyBytes, _ := GetKeyBytes([]byte(CUCCKey))

	// 传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错 可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	newPlainText, err := TripleDesDecrypt(cryptText, keyBytes, ivBytes...)
	if err != nil {
		return "", err
	}
	return string(newPlainText), nil
}

func TripleDesEncrypt(plainText ,key []byte, ivDes...byte)([]byte,error){
	if len(key)!= 24 {
		return nil, ErrKeyLengthTwentyFour
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil,err
	}
	paddingText := PKCS5Padding(plainText, block.BlockSize())

	var iv []byte
	if len(ivDes)!=0{
		if len(ivDes)!=8{
			return nil,ErrIvDes
		}else{
			iv=ivDes
		}
	}else{
		iv =[]byte(ivdes)
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte,len(paddingText))
	blockMode.CryptBlocks(cipherText,paddingText)
	return cipherText, nil
}

func TripleDesDecrypt(cipherText ,key []byte,ivDes...byte) ([]byte,error){
	if len(key)!=24{
		return nil, ErrKeyLengthTwentyFour
	}
	// 1. Specifies that the 3des decryption algorithm creates and returns a cipher.Block interface using the TDEA algorithm。
	block, err := des.NewTripleDESCipher(key)
	if err!=nil{
		return nil,err
	}

	// 2. Delete the filling
	// Before deleting, prevent the user from entering different keys twice and causing panic, so do an error handling
	defer func(){
		if err:=recover();err!=nil{
			switch err.(type){
			case runtime.Error:
				log.Println("runtime error:",err,"Check that the key is correct")
			default:
				log.Println("error:",err)
			}
		}
	}()

	var iv []byte
	if len(ivDes)!=0{
		if len(ivDes)!=8{
			return nil,ErrIvDes
		}else{
			iv=ivDes
		}
	}else{
		iv =[]byte(ivdes)
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)

	paddingText := make([]byte,len(cipherText)) //
	blockMode.CryptBlocks(paddingText,cipherText)


	plainText ,err:= PKCS5UnPadding(paddingText)
	if err!=nil{
		return nil,err
	}
	return plainText,nil
}


/**
1. Group plaintext
	If the blocksize is not an integer multiple of blocksize, the blocksize bit should be considered
    If des algorithm is used, the block size is 8 bytes
	With the AES algorithm, 16 bytes of fast size are filled in
    A tool for populating data when using block encryption mode
*/

// It is populated using pkcs5

func PKCS5Padding(plainText []byte, blockSize int) []byte{
	padding := blockSize - (len(plainText)%blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

func PKCS5UnPadding(plainText []byte)([]byte,error){
	length := len(plainText)
	number:= int(plainText[length-1])
	if number>=length{
		return nil,ErrPaddingSize
	}
	return plainText[:length-number],nil
}

/**
* desc : 联通密钥混淆
 */
func GetKeyBytes(keyBytes []byte) ([]byte, error){
	hasher := md5.New()
	hasher.Write(keyBytes)
	md5KeyBytes := hasher.Sum(nil)
	md5KeyBytes24 := make([]byte, 24)
	for i := 0; i < 16; i++ {
		md5KeyBytes24[i] = md5KeyBytes[i]
	}

	for i := 0; i < 8; i++ {
		md5KeyBytes24[i+16] = md5KeyBytes[i]
	}

	return md5KeyBytes24, nil
}
