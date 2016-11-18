package log2

import (
	l4g "github.com/alecthomas/log4go"
	"time"
)

const (
	filename = "rabbitgo.log"
)

func NewLogger() l4g.Logger {
	log := l4g.NewLogger()
	InitLogger(log)
	return log
}

func InitLogger(log l4g.Logger) {
	log.AddFilter("stdout", l4g.FINEST, l4g.NewConsoleLogWriter())

	flw := l4g.NewFileLogWriter(filename, false)
	flw.SetFormat("[%D %T] [%L] (%S) %M")
	flw.SetRotateSize(5000)
	flw.SetRotateLines(50)
	flw.SetRotateDaily(true)

	log.AddFilter("file", l4g.FINEST, flw)
}

func main() {

	log := NewLogger()
	//InitLogger(log)
	defer log.Close()

	log.Finest("Everything is created now (notice that I will not be printing to the file)")
	log.Fine("aaa")
	log.Debug("debug")
	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	log.Warn("warn")
	log.Critical("Time to close out!")

}
