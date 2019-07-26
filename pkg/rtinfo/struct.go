package rtinfo

import (
	"reflect"
)

type StructInfo struct {
	NameWithPkg string
}

func GetStructInfo(s interface{}) StructInfo {
	name := reflect.TypeOf(s).String()
	if reflect.ValueOf(s).Type().Kind() == reflect.Ptr {
		name = name[1:]
	}
	return StructInfo{
		NameWithPkg: name,
	}
}
