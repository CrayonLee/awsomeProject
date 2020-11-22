package main

import (
	"path"
	"runtime"
)

func GetCallerInfo(skip int) (fileName string,line int,funcName string) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok{
		return
	}
	fileName=path.Base(fileName)
	//根据pc拿到函数名
	funcName = runtime.FuncForPC(pc).Name()
	funcName =path.Base(funcName)
	return
}
