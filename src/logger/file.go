package main

import (
	"fmt"
	"os"
	"path"
	"time"
)


//文件日志结构体
//在这个版本中加入按文件大小切分日志的功能
type FileLogger struct {
	fileName string
	filePath string
	file     *os.File
	errFile  *os.File
	level    Level
	maxSize int64
}

//构造函数
func NewFielLogger(fileName, filePath string, level Level) *FileLogger {
	flo := &FileLogger{
		fileName: fileName,
		filePath: filePath,
		level:    level,
		//初始最大容量为10M
		maxSize: 10*1024*1024,
	}
	flo.initFile()
	return flo
}

//检查日志文件是否需要被切分
func(f *FileLogger) checkSplit(file *os.File) bool{
	fileInfo,_ := file.Stat()
	fileSize := fileInfo.Size()
	//当日志文件大于或等于maxSize时返回true
	return fileSize>=f.maxSize
}
//封装一个切分日志文件的方法
func (f *FileLogger) splitLogFile(file *os.File)  *os.File{

		//切分文件
		fileName:=file.Name()
		//定义一个文件的新名字
		backupName :=fmt.Sprintf("%s_%v.bak",fileName,time.Now().Unix())
		//1.把原来的文件关闭
		file.Close()
		//2.备份原来的文件
		os.Rename(fileName,backupName)
		//3.新建一个文件
		openFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err!=nil{
			panic(fmt.Errorf("打开日志文件失败"))
		}
		return openFile

}

//将指定日志文件打开  赋值给结构体
func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)
	fmt.Printf("fileName:%s\n", logName)
	//打开文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败,%v", logName, err))
	}
	f.file = fileObj

	//打开错误日志文件
	errLogName := fmt.Sprintf("%s.err", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败,%v", errLogName, err))
	}
	f.errFile = errFileObj
}

func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	msg := fmt.Sprintf(format, args)
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := GetCallerInfo(2)
	levelStr := GetLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", nowStr, fileName, line, funcName, levelStr, msg)

	//在往文件中写东西之前先做一个判断
	if f.checkSplit(f.file){
		f.file=f.splitLogFile(f.file)
	}

	fmt.Fprintln(f.file, logMsg)

	if level>=ErrorLevel {
		if f.checkSplit(f.errFile){
			f.errFile=f.splitLogFile(f.errFile)
		}

		fmt.Fprintln(f.errFile,logMsg)
	}
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DeubgLevel, format, args...)

}
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)

}
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)

}
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)

}
func (f *FileLogger) Fetal(format string, args ...interface{}) {
	f.log(FatalLevel, format, args...)

}

func(f *FileLogger) Close(){
	f.file.Close()
	f.errFile.Close()
}
