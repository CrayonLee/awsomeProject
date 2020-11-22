package main

var logger Logger
func main() {
	logger := NewFielLogger("a.log", "./", DeubgLevel)
	//logger:=NewConsoleLogger(InfoLevel)
	defer logger.Close()
	for{
		logger.Debug("这是一条Debug日志")
		logger.Info("这是一条Info日志")
		logger.Warn("这是一条Warn日志")
		logger.Error("这是一条错误日志")
		logger.Fetal("这是一条Feta日志")
	}




}