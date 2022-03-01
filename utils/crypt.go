package utils

import (
	"baseApiServer/models"
	"crypto/aes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// https://blog.csdn.net/laputa73/article/details/80476005
// https://tech1024.com/original/3015
// http://tool.chacuo.net/cryptaes
// https://www.keylala.cn/aes

func AESEncrypt(data, key string, paddingFunc func([]byte, int) []byte) (ret string, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("AesEncrypt: internal error! err=%v", p)
		}
	}()

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}

	content := paddingFunc([]byte(data), block.BlockSize())

	blockMode := NewECBEncryptor(block)
	crypt := make([]byte, len(content))
	blockMode.CryptBlocks(crypt, content)

	ret = strings.ToLower(hex.EncodeToString(crypt))
	return
}
func AESDecrypt(data, key string, unPaddingFunc func([]byte) []byte) (ret string, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("AesDecrypt: internal error! err=%v", p)
		}
	}()

	crypt, err := hex.DecodeString(strings.ToLower(data))
	if err != nil {
		return
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}

	blockMode := NewECBDecryptor(block)
	content := make([]byte, len(crypt))
	blockMode.CryptBlocks(content, crypt)

	content = unPaddingFunc(content)

	ret = string(content)
	return
}
func MakeNonCryptResponse(title string, response interface{}, unescape bool) (string, error) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		logObj.LogError("MakeNonCryptResponse", "Error", title, err.Error(), "", GetTimeToString(time.Now()))
		return ErrorMsg(models.ErrGetNonCryptRes), err
	}
	responseString := string(responseJSON)
	if unescape == true {
		responseString = strings.Replace(responseString, "\\u003c", "<", -1)
		responseString = strings.Replace(responseString, "\\u003e", ">", -1)
		responseString = strings.Replace(responseString, "\\u0026", "&", -1)
	}
	return responseString, nil
}
func MakeResponseEncryption(title string, encryptKey string, response interface{}, unescape bool) (string, error) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		logObj.LogError("MakeResponseEncryption", "Error", title, err.Error(), "", GetTimeToString(time.Now()))
		return ErrorMsg(models.ErrGetEncryptRes), err
	}
	responseString := string(responseJSON)
	if unescape == true {
		responseString = strings.Replace(responseString, "\\u003c", "<", -1)
		responseString = strings.Replace(responseString, "\\u003e", ">", -1)
		responseString = strings.Replace(responseString, "\\u0026", "&", -1)
	}
	responseEncrypted := responseString
	if encryptKey != "" {
		responseEncrypted, err = AESEncrypt(responseEncrypted, encryptKey, PKCS7Padding)
		if err != nil {
			logObj.LogError("MakeResponseEncryption", "Error", title, err.Error(), fmt.Sprintf("AESEncrypt failed! encryptKey=%s, response=%s\n", encryptKey, responseString), GetTimeToString(time.Now()))
			return ErrorMsg(models.ErrGetEncryptRes), err
		}
	}
	return responseEncrypted, nil
}
