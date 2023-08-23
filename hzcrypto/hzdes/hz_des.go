package hzdes

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
)

type encodingType int

const (
	defaultKey                   = "18689981"
	EncodingTypeSTD encodingType = 1
	EncodingTypeURL encodingType = 2
)

// Encrypt DES加密
func Encrypt(src string, encodingType encodingType, pwd ...string) (string, error) {
	key := defaultKey
	if len(pwd) > 0 {
		key = pwd[0]
	}
	s, err := desEncrypt([]byte(src), []byte(key))
	if err != nil {
		return "", err
	}

	var encode string
	if encodingType == EncodingTypeSTD {
		encode = base64.StdEncoding.EncodeToString(s)
	} else {
		encode = base64.URLEncoding.EncodeToString(s)
	}
	return encode, nil
}

// Decrypt DES解密接口
func Decrypt(src string, encodingType encodingType, pwd ...string) (string, error) {
	var decode []byte
	var err error
	if encodingType == EncodingTypeSTD {
		decode, err = base64.StdEncoding.DecodeString(src)
	} else {
		decode, err = base64.URLEncoding.DecodeString(src)
	}
	if err != nil {
		return "", err
	}

	key := defaultKey
	if len(pwd) > 0 {
		key = pwd[0]
	}
	s, err := desDecrypt(decode, []byte(key))
	if err != nil {
		return "", err
	}

	// 将结果拆分成str[]
	//var strAry = strings.Split(string(s), "|")

	return string(s), nil
}

func desEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()
	src = _PKCS5Padding(src, bs)
	if len(src)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}

	out := make([]byte, len(src))
	dst := out

	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func desDecrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}

	out = _PKCS5UnPadding(out)
	return out, nil
}

func _PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func _PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
