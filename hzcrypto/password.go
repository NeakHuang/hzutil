// Date: 2023/3/2
// Author:
// Description：

package hzcrypto

import (
	"github.com/NeakHuang/hzutil/hzcrypto/hzmd5"
)

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
