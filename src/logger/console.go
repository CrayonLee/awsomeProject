package main

import (
	"fmt"
	
	"time"
)

type ConsoleLogger struct {
	level Level
}


//构造函数
func NewConsoleLogger(level Level) *ConsoleLogger {
	return &ConsoleLogger{
		level:    level,
	}

}



func (f *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	msg := fmt.Sprintf(format, args)
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := GetCallerInfo(3)
	levelStr := GetLevelStr(level)
	fmt.Printf("[%s][%s:%d][%s][%s]%s\n", nowStr, fileName, line, funcName, levelStr, msg)

}

func (f *ConsoleLogger) Debug(format string, args ...interface{}) {
	f.log(DeubgLevel, format, args...)

}
func (f *ConsoleLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)

}
func (f *ConsoleLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)

}
func (f *ConsoleLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)

}
func (f *ConsoleLogger) Fetal(format string, args ...interface{}) {
	f.log(FatalLevel, format, args...)

}

func (f *ConsoleLogger) Close(){

}

