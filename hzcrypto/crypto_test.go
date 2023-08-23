package hzcrypto

import (
	"fmt"
	"github.com/NeakHuang/hzutil/hzcrypto/hzaes"
	"github.com/NeakHuang/hzutil/hzcrypto/hzmd5"
	"strings"
	"testing"
)

const secretSalt = "19880515"

func TestDesAndMd5(t *testing.T) {
	// 请求，添加钱包
	type addWalletFromPriReq struct {
		PrivateKey string `json:"privateKey"`
		SecretKey  string `json:"secretKey"`
		Index      int64  `json:"index"`
	}
	cPriKey := "82945a977b096c129f126c3e3292b69b4ad23de619d4d69c1e0c27559c9d8888"
	cSalt := "HipHop"
	cSecret := "BA7A63394C34886F"
	uid := "1"
	fmt.Println("原始：", cPriKey, cSalt)

	// 前端step1：对私钥和盐分别加密
	cAesVal, err := hzaes.Encrypt(cPriKey, uid)
	cAesSalt, err := hzaes.Encrypt(cSalt, uid)
	fmt.Println("Aes加密后Value：", cAesVal)
	fmt.Println("Aes加密后Salt：", cAesSalt)

	// 前端step2：对secretKey md5加密
	md5key := fmt.Sprintf("%s:%s", cSecret, uid)
	fmt.Println("Md5原始值：", md5key)
	cMD5SecretKey, err := hzmd5.EncryptString(md5key)
	if err != nil {
		fmt.Println("加密盐，解密失败", cMD5SecretKey)
		return
	}
	fmt.Println("Md5混淆值：", cMD5SecretKey)

	// 前端step3：将私钥和
	aesMixedPrivateKey, err := hzaes.Encrypt(fmt.Sprintf("%s|%s", cAesVal, cAesSalt), cMD5SecretKey)
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
	result, err := hzaes.Decrypt(privateKeyCrypto, sMD5SecretKey)
	resultList := strings.Split(string(result), "|")
	if len(resultList) >= 2 {
		aesPrivateKey, aseSalt := resultList[0], resultList[1]
		privateKey, _ := hzaes.Decrypt(aesPrivateKey, uid)
		salt, _ := hzaes.Decrypt(aseSalt, uid)
		fmt.Println(string(privateKey), string(salt))
	}

}
