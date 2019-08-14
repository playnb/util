package log

import "testing"

func TestDefaultLogger(t *testing.T) {
	Init(DefaultLogger(`C:\ltp\code\log`, "test"))
	InitPanic(`C:\ltp\code\panic`)
	defer Flush()

	Debug("debug........")
	Info("info........")
	Trace("trace........")
	Error("error........")
	Fatal("fatal........")

	defer PrintPanicStack()
	MakePanic()
}
