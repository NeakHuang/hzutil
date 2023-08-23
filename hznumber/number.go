// Date: 2023/2/28
// Author:
// Description：

package hznumber

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Numeric interface {
	Int | Uint | Float
}

// type SliceNumeric[T Numeric] []T

const (
	SizeKB = 1024
	SizeMB = 1024 * SizeKB

	Uint64Max = ^uint64(0) // = uint64(math.MaxUint64)
)

func BoolToNumber[T Numeric](ok bool) T {
	if ok {
		return 1
	}
	return 0
}
