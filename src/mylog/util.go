package mylog

import (
	"path"
	"runtime"
)

func getCallerInfo() (fileName,funcName string,line int)  {
	pc,fileName,line,ok := runtime.Caller(3)
	if !ok{
		return
	}
	funcName =runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	fileName = path.Base(fileName)
	//fmt.Println(funcName,fileName,line)
	return
}