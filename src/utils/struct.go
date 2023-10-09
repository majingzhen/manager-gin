package utils

import "reflect"

// GetFieldValues 获取结构体切片中某个字段的值
func GetFieldValues(structSlice interface{}, fieldName string) []interface{} {
	var values []interface{}
	structs := structSlice.([]interface{})
	for _, structData := range structs {
		values = append(values, reflect.ValueOf(structData).FieldByName(fieldName).Interface())
	}
	return values
}
