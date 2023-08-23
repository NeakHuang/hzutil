/**
 *	Date: 2019/5/27
 *  Desc:
 *	extra:
 */
package encrypt

import (
	"crypto/rc4"
	"fmt"
	"testing"
)

func TestRC4(t *testing.T) {
	key := []byte("fund_test")
	str := []byte("ZN5I6DGWELOIBN2RKVR77NHP77MT76PO")
	// 加密操作
	dest1 := make([]byte, len(str))
	fmt.Printf("方法1加密前:%s \n", str)
	cipher1, _ := rc4.NewCipher(key)
	cipher1.XORKeyStream(dest1, str)
	fmt.Printf("方法1加密后:%s \n", dest1)

	// 解密操作
	dest2 := make([]byte, len(dest1))
	cipher2, _ := rc4.NewCipher(key) // 切记：这里不能重用cipher1，必须重新生成新的
	cipher2.XORKeyStream(dest2, dest1)
	fmt.Printf("方法1解密后:%s \n\n", dest2)
}

func TestDES(t *testing.T) {
	// key := []byte("12345678download87654321somebody")
	// str := []byte("ZN5I6DGWELOIBN2RKVR77NHP77MT76PO")
	// // 加密操作
	// fmt.Printf("方法1加密前:%s \n", str)
	// dest1, err := gaes.Encrypt(str, key)
	// fmt.Println(err)
	// fmt.Printf("方法1加密后:%s \n", dest1)
	// dest1Str := gbase64.EncodeToString(dest1)
	// fmt.Printf("方法1 base64加密后:%s \n", dest1Str)
	//
	// // 解密操作
	// dest2Str, _ := gbase64.Decode([]byte(dest1Str))
	// fmt.Printf("方法1 base64解密后:%s \n", dest2Str)
	// dest2, _ := gaes.Decrypt(dest2Str, key) // 切记：这里不能重用cipher1，必须重新生成新的
	// fmt.Printf("方法1解密后:%s \n\n", dest2)
}
