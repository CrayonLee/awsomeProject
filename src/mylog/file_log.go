package mylog

import (
	"fmt"
	"os"
	"time"
)

//定义一个往文件中记录日志的结构体
type FileLogger struct{
	level int
	logFilePath string
	logFileName string
	logFile *os.File
}

//专门用来初始化文件日志的文件句柄
func (f *FileLogger) initFileLogger()  {
	filePath := fmt.Sprintf("%s/%s",f.logFilePath,f.logFileName)
	file,err := os.OpenFile(filePath,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err!=nil{
		panic(fmt.Sprintf("open file: %s failed",filePath))
	}
	f.logFile=file

}

//构造方法
func NewFileLogger(level int,logFilePath,logFileName string) *FileLogger {
	flObj := &FileLogger{
		level: level,
		logFilePath: logFilePath,
		logFileName: logFileName,
	}
	flObj.initFileLogger()
	return flObj
}

func (f *FileLogger) writeLog(msg string,args...interface{})  {
	//丰富日志格式
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	fileName, funcName, line := getCallerInfo()
	msg = fmt.Sprintf("%s [%s] [%s-%s] [%d] %s",nowStr,getLevelStr(f.level),fileName,funcName,line,msg)
	//f.logFile.WriteString(msg)
	fmt.Fprintf(f.logFile,msg,args...)
	fmt.Fprintln(f.logFile)	//加换行
}
//记录日志
func (f *FileLogger) Debug(msg string,args...interface{})  {

	f.writeLog(msg)
}

func (f *FileLogger) Trace(msg string,args...interface{}){
	f.writeLog(msg)
}

func (f *FileLogger) Info(msg string,args...interface{}){
	f.writeLog(msg)
}

func (f *FileLogger) Warn(msg string,args...interface{}){
	f.writeLog(msg)
}
func (f *FileLogger) Error(msg string,args...interface{}){
	f.writeLog(msg)
}
func (f *FileLogger) Critical(msg string,args...interface{}){
	f.writeLog(msg)
}






