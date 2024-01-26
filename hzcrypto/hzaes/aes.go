package hzaes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

// aes加密
func Encrypt(str, pwd string) (encode string, err error) {
	var key = []byte(pwd)
	md5sum := md5.Sum([]byte(pwd))
	key = []byte(fmt.Sprintf("%x", md5sum))
	//key = md5sum[:] // 不再直接截取
	result, err := AesEncrypt([]byte(str), key)
	if err != nil {
		return
	}
	encode = base64.StdEncoding.EncodeToString(result)
	//fmt.Printf("strsdata = %v  \n", hex.EncodeToString(result))
	return
}

// aes解密
func Decrypt(encode, pwd string) (str []byte, err error) {
	re, err1 := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		err = err1
		return
	}
	md5sum := md5.Sum([]byte(pwd))
	key := []byte(fmt.Sprintf("%x", md5sum))
	//key = md5sum[:] // 不再直接截取
	str, err = AesDecrypt(re, key)
	return
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

// 这个方案必须和js的方法是一样的
func PaddingLeft(ori []byte, pad byte, length int) []byte {
	if len(ori) >= length {
		return ori[:length]
	}
	pads := bytes.Repeat([]byte{pad}, length-len(ori))
	return append(pads, ori...)
}
