// Date: 2023/3/2
// Author:
// Description：

package hzpwd

import (
	"crypto/sha256"
	"fmt"
	"github.com/NeakHuang/hzutil/hzcrypto/hzmd5"
)

var (
	PasswordPower = 1
)

func SetPasswordPower(power int) {
	PasswordPower = power
}

// PasswordMD5  encrypts password using MD5 algorithms.
// Salt is for confusion
func PasswordMD5(password string, salt ...string) string {
	if len(password) <= 0 {
		return ""
	}
	pwd := hzmd5.MustEncryptString(password)
	// using salt for confusion
	if len(salt) > 0 {
		pwd += hzmd5.MustEncryptString(salt[0])
	}
	return hzmd5.MustEncryptString(pwd)
}

func Sha256ByString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	arr := h.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func Sha256ByStringWithCount(s string, c int) string {
	v := s
	for i := 0; i < c; i++ {
		v = Sha256ByString(v)
	}
	return v
}

func EncodePassword(s string) string {
	return Sha256ByStringWithCount(s, PasswordPower)
}
