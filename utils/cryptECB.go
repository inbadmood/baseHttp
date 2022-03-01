package utils

import "crypto/cipher"

// it seems imitating from crypto\cipher\cbc.go
type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncryptor ecb

func NewECBEncryptor(b cipher.Block) cipher.BlockMode {
	return (*ecbEncryptor)(newECB(b))
}
func (x *ecbEncryptor) BlockSize() int { return x.blockSize }
func (x *ecbEncryptor) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecryptor ecb

func NewECBDecryptor(b cipher.Block) cipher.BlockMode {
	return (*ecbDecryptor)(newECB(b))
}
func (x *ecbDecryptor) BlockSize() int { return x.blockSize }
func (x *ecbDecryptor) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
