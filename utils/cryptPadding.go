package utils

import (
	"bytes"
)

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	// reference of pad_zero at
	// https://www.php.net/manual/en/function.openssl-encrypt.php
	if len(ciphertext)%blockSize == 0 {
		return ciphertext
	}
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padText...)
}
func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	// So fundamentally PKCS#5 padding is a subset of PKCS#7 padding for 8 byte block sizes. Hence, PKCS#5 padding can not be used for AES. PKCS#5 padding was only defined with RC2/RC5 and (triple) DES operation in mind.
	// https://crypto.stackexchange.com/questions/9043/what-is-the-difference-between-pkcs5-padding-and-pkcs7-padding
	// https://www.twblogs.net/a/5d193e52bd9eee1e5c82dfe1
	return PKCS7Padding(ciphertext, blockSize)
}
func PKCS5UnPadding(origData []byte) []byte {
	return PKCS7UnPadding(origData)
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	// reference of pkcs5_pad at
	// https://my.oschina.net/luoxiaojun1992/blog/883123
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	if length == 0 {
		return origData
	}
	unPadding := int(origData[length-1])
	if length < unPadding {
		panic("unPadding error")
	}
	return origData[:(length - unPadding)]
}
