// Date: 2023/5/8
// Author:
// Description：

package hzreflect

import (
	"github.com/NeakHuang/hzutil/hzstr"
	"reflect"
)

// ReflectField
// pointData: reference, 指针传递参数
// valueData: value, 值传递参数
func ReflectField(pointData any, valueData any, tagKey ...string) any {
	typ := reflect.TypeOf(valueData)
	fieldNum := typ.NumField()

	vo := reflect.ValueOf(pointData).Elem()
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

	// fn := func(rt reflect.Type, vo reflect.Value) (list []string) {
	// 	fieldNum := rt.NumField()
	// 	for i := 0; i < fieldNum; i++ {
	// 		key := rt.Field(i).Name
	// 		vo.FieldByName(key).SetString(key) // 设置值
	// 		list = append(list, key)
	// 	}
	// 	return list
	// }
	// data := new(Data)
	// typ, val := reflect.TypeOf(*data), reflect.ValueOf(data).Elem()
	// fn(typ, val)
}
