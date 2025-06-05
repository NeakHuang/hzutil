// Package hzslice
// @Author Neak
// @Date 周四 25/06/05
// @Description
package hzslice

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	// 示例用法
	str := "1,2,3,4,5"

	// 转换为int切片
	intSlice, err := StrToNumSlice[int](str, ",")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("int slice: %v (type: %T)\n", intSlice, intSlice)
	}

	// 转换为float64切片
	floatSlice, err := StrToNumSlice[float64](str, ",")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("float64 slice: %v (type: %T)\n", floatSlice, floatSlice)
	}

	// 转换为uint32切片
	uintSlice, err := StrToNumSlice[uint32](str, ",")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("uint32 slice: %v (type: %T)\n", uintSlice, uintSlice)
	}

	// 尝试不支持的类型
	_, err = StrToNumSlice[bool](str, ",")
	if err != nil {
		fmt.Println("Expected error for bool:", err)
	}
}

func TestConvertExample(t *testing.T) {
	// 转换为[]int
	intSlice, _ := StrToNumSlice[int]("1,2,3", ",")
	fmt.Printf("int slice: %v (type: %T)\n", intSlice, intSlice)
	// 转换为[]float64
	floatSlice, _ := StrToNumSlice[float64]("1.1|2.2|3.3", "|")
	fmt.Printf("float64 slice: %v (type: %T)\n", floatSlice, floatSlice)
	// 转换为[]uint32
	uintSlice, _ := StrToNumSlice[uint32]("10 20 30", " ")
	fmt.Printf("uint32 slice: %v (type: %T)\n", uintSlice, uintSlice)
}

func TestSliceToString(t *testing.T) {
	intSlice := []int{1, 2, 3, 4}
	fmt.Println(SliceToString(intSlice, ",")) // 输出: 1,2,3,4

	strSlice := []string{"a", "b", "c"}
	fmt.Println(SliceToString(strSlice, "|")) // 输出: a|b|c

	floatSlice := []float64{1.1, 2.2, 3.3}
	fmt.Println(SliceToString(floatSlice, " ")) // 输出: 1.1 2.2 3.3
}
