package log

import (
	"fmt"
	"github.com/cihub/seelog"
	"os"
	"runtime"
)

type see struct {
	logger seelog.LoggerInterface
}

func (s *see) Debug(format string, a ...interface{}) {
	s.logger.Debug(fmt.Sprintf(format, a...))
}
func (s *see) Info(format string, a ...interface{}) {
	s.logger.Info(fmt.Sprintf(format, a...))
}
func (s *see) Trace(format string, a ...interface{}) {
	s.logger.Trace(fmt.Sprintf(format, a...))
}
func (s *see) Error(format string, a ...interface{}) {
	s.logger.Error(fmt.Sprintf(format, a...))
}
func (s *see) Fatal(format string, a ...interface{}) {
	s.logger.Critical(fmt.Sprintf(format, a...))
}
func (s *see) Warn(format string, a ...interface{}) {
	s.logger.Warn(fmt.Sprintf(format, a...))
}
func (s *see) Flush() {
	s.logger.Flush()
}

func DefaultLogger(logDir string, logName string) logger {
	if len(logDir) == 0 && len(logName) == 0 {
		sl := &see{}
		sl.logger = seelog.Default
		return sl
	}
	os.MkdirAll(logDir, os.ModePerm)

	sl := &see{}
	logConfig := `
<seelog>
    <outputs formatid="main">
		<filter levels="info,critical,error,debug,trace,warn">
		`
	//只有windows
	if runtime.GOOS == "windows" {
		logConfig += `<console />`
	}
	logConfig += `
			<rollingfile type="date" filename="` + logDir + `/` + fmt.Sprintf("%s.log", logName) + `" datepattern="2006.01.02-15" />
        </filter>
    </outputs>

    <formats>
        <format id="main" format="%Date %Time [%LEV] [%File:%Line] %Msg%n"/>
    </formats>
</seelog>
`

	sl.logger, _ = seelog.LoggerFromConfigAsBytes([]byte(logConfig))
	sl.logger.SetAdditionalStackDepth(3)
	return sl
}
