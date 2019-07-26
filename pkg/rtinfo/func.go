package rtinfo

import (
	"reflect"
	"runtime"
)

type FuncInfo struct {
	NameWithPkg string
	FileName    string
	LineNumber  int
}

func GetFuncInfo(f interface{}) FuncInfo {
	_, filename, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return FuncInfo{
		NameWithPkg: name,
		FileName:    filename,
		LineNumber:  line,
	}
}
