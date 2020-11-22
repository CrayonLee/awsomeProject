package main

type Level uint16

//定义具体日志级别changliang
const (
	DeubgLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

//定义一个Logger接口
type Logger interface {
	Debug( string,  ...interface{})
	Info( string,  ...interface{})
	Warn( string,  ...interface{})
	Error( string,  ...interface{})
	Fetal( string,  ...interface{})
	Close()
}

func GetLevelStr(level Level) string{
	switch level {
	case DeubgLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FETAL"
	default:
		return "DEBUG"
	}
}


