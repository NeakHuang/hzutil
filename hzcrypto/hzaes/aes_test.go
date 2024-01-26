package hzaes

import (
	"fmt"
	"github.com/NeakHuang/hzutil/hzcrypto/hzmd5"
	"strings"
	"testing"
)

func TestAESCrypto(t *testing.T) {
	// AES加密需要一个16字节的密钥
	//key := []byte("a very very very very secret key")
	key := []byte("191b62e5947d12d4")

	// 需要加密的明文
	plaintext := []byte("some really really really long plaintext")

	ciphertext, err := AesEncrypt(key, plaintext)
	fmt.Println(fmt.Sprintf("%x", ciphertext))
	if err != nil {
		panic(err)
	}

	decryptedText, err := AesDecrypt(key, ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%s", decryptedText))
}

func TestDesAndMd5(t *testing.T) {
	// 请求，添加钱包
	type addWalletFromPriReq struct {
		PrivateKey string `json:"privateKey"`
		SecretKey  string `json:"secretKey"`
		Index      int64  `json:"index"`
	}
	cPriKey := "1234c0b9b8ea54c820402b305479c39e007b37989e9124f35b1e8763b07ab123"
	cSalt := "TestCD"
	cSecret := "AAAA63366C348888"
	uid := "1"
	fmt.Println("原始：", cPriKey, cSalt)

	// 前端step1：对私钥和盐分别加密
	cAesPriKey, err := Encrypt(cPriKey, uid)
	if err != nil {
		fmt.Println(err)
		return
	}
	cAesSalt, err := Encrypt(cSalt, uid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Aes加密后Value：", cAesPriKey)
	fmt.Println("Aes加密后Salt：", cAesSalt)

	// 前端step2：对secretKey md5加密
	md5key := fmt.Sprintf("%s:%s", cSecret, uid)
	fmt.Println("Md5原始值：", md5key)
	cMD5SecretKey, err := hzmd5.EncryptString(md5key)
	if err != nil {
		fmt.Println("加密盐，解密失败", err)
		return
	}
	fmt.Println("Md5混淆值：", cMD5SecretKey)

	// 前端step3：将私钥和
	aesMixedPrivateKey, err := Encrypt(fmt.Sprintf("%s|%s", cAesPriKey, cAesSalt), cMD5SecretKey)
	if err != nil {
		fmt.Println("加密发送失败", err)
		return
	}
	fmt.Println(aesMixedPrivateKey)

	// 请求数据
	in := &addWalletFromPriReq{
		PrivateKey: aesMixedPrivateKey,
		SecretKey:  cSecret, // 加密key。
		Index:      0,       // 钱包index，用于获取钱包得公钥，默认0
	}
	fmt.Println(in.PrivateKey)
	// ——————————————————————————————————————————————————————————

	//后端：收到解密
	//后端step1：对secretKey 核对
	sMd5Key := fmt.Sprintf("%s:%s", in.SecretKey, uid)
	sMD5SecretKey, err := hzmd5.EncryptString(sMd5Key)
	if err != nil {
		fmt.Println("加密盐，解密失败", err)
		return
	}

	// 后端step2：对privateKey解密
	privateKeyCrypto := in.PrivateKey
	result, err := Decrypt(privateKeyCrypto, sMD5SecretKey)
	resultList := strings.Split(string(result), "|")
	if len(resultList) >= 2 {
		aesPrivateKey, aseSalt := resultList[0], resultList[1]
		privateKey, _ := Decrypt(aesPrivateKey, uid)
		salt, _ := Decrypt(aseSalt, uid)
		fmt.Println(string(privateKey), string(salt))
	}

}
