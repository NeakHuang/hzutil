// Package hzslice
// @Author Neak
// @Date 周四 25/06/05
// @Description
package hzslice

import (
	"fmt"
	"github.com/NeakHuang/hzutil/hznumber"
	"reflect"
	"strconv"
	"strings"
)

// NumSliceToStr 数字数组转string
func NumSliceToStr[T hznumber.Numeric](valList []T, split string) (str string) {
	return NumSliceToString(valList, split)
}

// NumSliceToString 数字数组转string
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

// StrToNumSlice 将字符串转换为任意数值类型的切片
func StrToNumSlice[T any](str, split string) ([]T, error) {
	if len(str) == 0 {
		return nil, nil
	}

	// 获取目标类型的kind
	targetType := reflect.TypeOf((*T)(nil)).Elem()
	var kind reflect.Kind
	if targetType != nil {
		kind = targetType.Kind()
	}

	// 支持的数值类型检查
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		// 这些都是支持的数值类型
	default:
		return nil, fmt.Errorf("unsupported numeric type: %v", targetType)
	}

	// 分割字符串
	strSlice := strings.Split(str, split)
	result := make([]T, 0, len(strSlice))

	// 逐个转换
	for _, s := range strSlice {
		if s == "" {
			continue
		}

		val, err := convertToType[T](s)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}

	return result, nil
}

// convertToType 将字符串转换为指定类型
func convertToType[T any](s string) (T, error) {
	var zero T
	targetType := reflect.TypeOf(zero)

	switch targetType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(s, 10, targetType.Bits())
		if err != nil {
			return zero, err
		}
		return reflect.ValueOf(val).Convert(targetType).Interface().(T), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(s, 10, targetType.Bits())
		if err != nil {
			return zero, err
		}
		return reflect.ValueOf(val).Convert(targetType).Interface().(T), nil

	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(s, targetType.Bits())
		if err != nil {
			return zero, err
		}
		return reflect.ValueOf(val).Convert(targetType).Interface().(T), nil

	default:
		return zero, fmt.Errorf("unsupported type: %v", targetType)
	}
}

// SliceToString 将任意slice转换为字符串，使用指定分隔符
func SliceToString[T any](slice []T, sep string) string {
	if len(slice) == 0 {
		return ""
	}

	// 使用strings.Builder高效拼接字符串
	var builder strings.Builder
	for i, v := range slice {
		if i > 0 {
			builder.WriteString(sep)
		}
		builder.WriteString(fmt.Sprintf("%v", v))
	}
	return builder.String()
}
