package log2

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/xww/rabbitgo/conf"
	"time"
)

const (
	filename = "rabbitgo.log"
)

func NewLogger(config map[string]interface{}) l4g.Logger {
	log := l4g.NewLogger()
	InitLogger(log, config)
	return log
}

func InitLogger(log l4g.Logger, config map[string]interface{}) {
	log.AddFilter("stdout", l4g.FINEST, l4g.NewConsoleLogWriter())
	flw := l4g.NewFileLogWriter(filename, config["logEnableRotate"].(bool))
	flw.SetFormat(config["logFormat"].(string))
	flw.SetRotateSize(int(config["logRotateSize"].(int64)))
	flw.SetRotateLines(int(config["logRotateLines"].(int64)))
	flw.SetRotateDaily(config["logRotateDaily"].(bool))
	log.AddFilter("file", l4g.FINEST, flw)
}

func main2() {

	log := NewLogger(conf.InitConfig())

	//InitLogger(log)
	defer log.Close()

	log.Finest("Everything is created now (notice that I will not be printing to the file)")
	log.Fine("aaa")
	log.Debug("debug")
	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	log.Warn("warn")
	log.Critical("Time to close out!")

}
