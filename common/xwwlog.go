package common

import (
	"io"
	"log"
)

type XwwLog struct {
	l *log.Logger
}

func NewLog(out io.Writer) *XwwLog {
	return &XwwLog{l: log.New(out, "xxx", log.Lshortfile|log.LstdFlags)}
}

func (bl *XwwLog) Debug(s interface{}) {
	bl.l.SetPrefix("[Debug  ]")
	bl.l.Println(s)

}
func (bl *XwwLog) Info(s interface{}) {
	bl.l.SetPrefix("[Info   ]")
	bl.l.Println(s)
}
func (bl *XwwLog) Warning(s interface{}) {
	bl.l.SetPrefix("[Warning]")
	bl.l.Println(s)
}
func (bl *XwwLog) Error(s interface{}) {
	bl.l.SetPrefix("[Error  ]")
	bl.l.Println(s)
}
func (bl *XwwLog) Fatal(s interface{}) {
	bl.l.SetPrefix("[Fatal  ]")
	bl.l.Println(s)
}

func (bl *XwwLog) SetFormat(flags int) {
	bl.l.SetFlags(flags)
}

func (bl *XwwLog) Println(s interface{}) {
	bl.l.Println(s)
}

func AA() {
	println("aaa")
}

/*
func main(){

	fileName2 := "xww.log"
	logFile2,err2  := os.Create(fileName2)
	defer logFile2.Close()
	if err2 != nil {
		log.Fatalln("open file error !")
	}
	mylog := NewLog(logFile2)
	mylog.Debug("debug")
	mylog.Fatal("fatal")
	mylog.Info("info")


}*/
