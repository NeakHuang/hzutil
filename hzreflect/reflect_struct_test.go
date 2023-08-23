// Date: 2023/5/8
// Author:
// Description：

package hzreflect

import (
	"fmt"
	"github.com/NeakHuang/hzutil/hzstr"
	"reflect"
	"testing"
)

type (
	eString struct {
		A string
		B string
		C string
	}

	eNumber struct {
		One   int32  `tag:"1"`
		Two   int32  `tag:"2"`
		Three int32  `tag:"3"`
		Test  string `tag:"Test1"`
	}
)

func TestReflectField(t *testing.T) {

	eStr := ReflectField(&eString{}, eString{}).(*eString)
	fmt.Println(eStr.A, eStr.B, eStr.C)

	eNum := ReflectField(&eNumber{}, eNumber{}).(*eNumber)
	fmt.Println(eNum.One, eNum.Two, eNum.Three, eNum.Test)

	str := &eString{}
	num := &eNumber{}
	for point, value := range map[any]any{str: *str, num: *num} {
		reflectField(point, value)
	}
	fmt.Println(str.A, str.B, str.C)
	fmt.Println(num.One, num.Two, num.Three, num.Test)

}

func reflectField(pointData any, valueData any, tagKey ...string) any {
	typ := reflect.TypeOf(valueData)
	fieldNum := typ.NumField()

	pointVo := reflect.ValueOf(pointData)
	vo := pointVo.Elem()
	for i := 0; i < fieldNum; i++ {
		field := typ.Field(i)
		tag := "tag"
		if len(tagKey) > 0 {
			tag = tagKey[0]
		}
		if tagVal, ok := field.Tag.Lookup(tag); ok {
			kind := vo.FieldByName(field.Name).Kind()
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				val := hzstr.ToInt64(tagVal)
				vo.FieldByName(field.Name).SetInt(val) // 设置值
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				val := hzstr.ToUint64(tagVal)
				vo.FieldByName(field.Name).SetUint(val) // 设置值
			case reflect.String:
				vo.FieldByName(field.Name).SetString(tagVal) // 设置值
			}
		} else {
			key := field.Name
			vo.FieldByName(key).SetString(key) // 设置值
		}
		// name := vo.FieldByName(typ.Field(i).Name).String() // 获取值名称
		// fmt.Println(field.Name, name)
	}
	return pointData

}
