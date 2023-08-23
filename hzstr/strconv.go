// Date: 2023/2/17
// Author:
// Description：

package hzstr

import (
	"github.com/NeakHuang/hzutil/hznumber"
	"strconv"
	"strings"
)

func IntConv(v int) string {
	str := strconv.Itoa(v)
	return str
}

func UintConv(v uint) string {
	str := strconv.Itoa(int(v))
	return str
}

func Int32Conv(v int32) string {
	str := strconv.Itoa(int(v))
	return str
}

func Uint32Conv(v uint32) string {
	str := strconv.Itoa(int(v))
	return str
}

func Int64Conv(v int64) string {
	str := strconv.Itoa(int(v))
	return str
}

func Uint64Conv(v uint64) string {
	str := strconv.Itoa(int(v))
	return str
}

func Float64Conv(v float64) string {
	str := strconv.FormatFloat(v, 'e', -1, 64)
	return str
}

func ToBool(v string) bool {
	if v = strings.Trim(v, " "); v == "" {
		return false
	}
	b, _ := strconv.ParseBool(v)
	return b
}

func ToInt(v string) int {
	if v = strings.Trim(v, " "); v == "" {
		return 0
	}
	if i, err := strconv.ParseInt(v, 0, 32); err == nil {
		return int(i)
	} else {
		return 0
	}
}

func ToInt32(v string) int32 {
	if v = strings.Trim(v, " "); v == "" {
		return 0
	}
	if i, err := strconv.ParseInt(v, 0, 32); err == nil {
		return int32(i)
	} else {
		return 0
	}
}

func ToUint(v string) uint {
	if v = strings.Trim(v, " "); v == "" {
		return 0
	}
	if i, err := strconv.ParseInt(v, 0, 32); err == nil {
		return uint(i)
	} else {
		return 0
	}
}

func ToUint32(v string) uint32 {
	if v = strings.Trim(v, " "); v == "" {
		return 0
	}
	if i, err := strconv.ParseInt(v, 0, 32); err == nil {
		return uint32(i)
	} else {
		return 0
	}
}

func ToInt64(v string) int64 {
	if v = strings.Trim(v, " "); v == "" {
		return 0
	}
	if i, err := strconv.ParseInt(v, 0, 64); err == nil {
		return i
	} else {
		return 0
	}
}
func ToUint64(v string) uint64 {
	if v = strings.Trim(v, " "); v == "" {
		return 0
	}
	if i, err := strconv.ParseUint(v, 0, 64); err == nil {
		return uint64(i)
	} else {
		return 0
	}
}

func ToFloat(v string) float32 {
	return float32(ToFloat64(v))
}

func ToFloat64(v string) float64 {
	if v = strings.Trim(v, " "); v == "" {
		return 0
	}
	index := strings.LastIndex(v, ".")
	if index == -1 {
		return float64(ToInt(v))
	}
	if f, err := strconv.ParseFloat(v, len(v)-1-index); err == nil {
		return f
	} else {
		return 0
	}
}

// 数字数组转string
func NumSliceToString[T hznumber.Numeric](valList []T, split string) (str string) {
	if nil == valList || len(valList) == 0 {
		return str
	}
	if len(split) <= 0 {
		split = ","
	}
	for _, s := range valList {
		str += strconv.Itoa(int(s)) + split
	}
	cutLen := len(split)
	return str[0 : len(str)-cutLen]
}

// 字符数组转string
// 之后替换为 strings.Join
func StrSliceToString(valList []string, split string) (str string) {
	// if nil == valList || len(valList) == 0 {
	// 	return str
	// }
	// if len(split) <= 0 {
	// 	split = ","
	// }
	// for _, s := range valList {
	// 	str += s + split
	// }
	// cutLen := len(split)
	// return str[0 : len(str)-cutLen]
	return strings.Join(valList, split)
}
