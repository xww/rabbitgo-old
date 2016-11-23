package main

import (
	"github.com/judwhite/go-svc/svc"
	"syscall"
	"fmt"
	"path/filepath"
	"os"

	"sync"
	"time"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}

type context struct {
	a *A
}

type A struct {
	vara string
	waitGroup    WaitGroupWrapper
}
func New() *A{
	a := &A{
		vara:"aa",
	}
	return a
}
func myfunc() {
	for{
		timestamp := time.Now().Unix()
		tm := time.Unix(timestamp, 0)
		fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
		time.Sleep(2 * time.Second)
	}

}
func (a *A) Main() {
	//ctx := &context{a}

	a.waitGroup.Wrap(func() {
		myfunc()
	})
}
func (a *A) Exit() {

	a.waitGroup.Wait()
}

type program struct {
	 a *A
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		fmt.Print("error")
	}
}

func (p *program) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

func (p *program) Start() error {
	myaa := New()
	myaa.Main()
	return  nil

}

func (p *program) Stop() error {
	if p.a != nil {
		p.a.Exit()
	}
	return nil
}


