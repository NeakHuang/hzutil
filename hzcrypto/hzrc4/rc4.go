/**
 *	Date: 2019/5/27
 *  Desc:
 *	extra:
 */
package encrypt

import (
	"math/rand"
	"time"
)

type RC4 struct {
	bytes []byte
	i     int
	j     int
}

func NewRC4() *RC4 {
	return &RC4{
		bytes: make([]byte, 256, 256),
	}
}

func (r *RC4) Init(key []byte) {
	var keyLen = len(key)
	var T = make([]byte, 256, 256)
	for i := 0; i < 256; i++ {
		r.bytes[i] = byte(i)
		T[i] = key[i%keyLen]
	}

	var j = 0
	for i := 0; i < 256; i++ {
		j = j + int(r.bytes[i]) + int(T[i])
		j = (j % 256) & 0xFF
		Swap(r.bytes, i, j)
	}
}

func (r *RC4) Crypt(data []byte, start int, end int) {
	var bytes = r.bytes
	if end <= 0 {
		end = len(data)
	}
	for i := start; i < end; i++ {
		r.i += 1
		r.i = (r.i % 256) & 0xFF

		r.j = r.j + int(bytes[r.i])
		r.j = (r.j % 256) & 0xFF

		Swap(bytes, r.i, r.j)

		var t = int(bytes[r.i] + bytes[r.j])
		t = (t % 256) & 0xFF
		var b = bytes[t]
		data[i] = byte(b ^ data[i])
	}
}

func Swap(bytes []byte, i, j int) []byte {
	var t = bytes[i]
	bytes[i] = bytes[j]
	bytes[j] = t
	return bytes
}

func GetRandomKey() []byte {
	rand.Seed(time.Now().Unix()) // 种子到时候改成固定的一个约定数值
	var len = 10 + rand.Intn(22)
	var key = make([]byte, 256, 256)
	for i := 0; i < len; i++ {
		key[i] = byte(rand.Intn(256))
	}
	return key
}
