package log

type logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Flush()
}

var _ins logger

func Init(l logger) {
	_ins = l
}

func Debug(format string, a ...interface{}) {
	_ins.Debug(format, a...)
}

func Info(format string, a ...interface{}) {
	_ins.Info(format, a...)
}

func Trace(format string, a ...interface{}) {
	_ins.Trace(format, a...)
}

func Error(format string, a ...interface{}) {
	_ins.Error(format, a...)
}

func Fatal(format string, a ...interface{}) {
	_ins.Fatal(format, a...)
}

func Warn(format string, a ...interface{}) {
	_ins.Warn(format, a...)
}

func Flush() {
	_ins.Flush()
}
