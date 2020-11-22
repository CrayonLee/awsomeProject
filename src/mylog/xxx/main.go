package main

import "mylog"

func main() {
	logger := mylog.NewFileLogger(mylog.DEBUG, "", "abc.log")
	logger.Debug("这是一条Debug日志")
	logger= mylog.NewFileLogger(mylog.TRACE, "", "abc.log")
	logger.Trace("这是一条Trace日志")
	logger = mylog.NewFileLogger(mylog.WARN, "", "abc.log")
	logger.Warn("这是一条Warn日志")
	logger = mylog.NewFileLogger(mylog.ERROR, "", "abc.log")
	logger.Error("这是一条Error日志")
	logger = mylog.NewFileLogger(mylog.CRITICAL, "", "abc.log")
	logger.Critical("这是一条Critical日志")


}
