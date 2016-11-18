package log2

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/xww/rabbitgo/conf"
)

var Log l4g.Logger

func init() {
	config := conf.InitConfig()
	Log = NewLogger(config)
}
