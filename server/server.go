package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/xww/rabbitgo/conf"
	"github.com/xww/rabbitgo/log2"
	"log"
	"net"
	"time"
)

var (
	flagSet = flag.NewFlagSet("nsqadmin", flag.ExitOnError)
	config  = flagSet.String("config", "./rabbitgo.conf", "path to config file")
)

func handleConnection(conn net.Conn) {
	data, err := bufio.NewReader(conn).ReadString('\n')
	conn.RemoteAddr().Network()
	if err != nil {
		log.Fatal("get client data error: ", err)
	}
	fmt.Printf("%#v\n", data)
	fmt.Fprintf(conn, "hello client\n")
	conn.Close()
}

func main() {
	appConfig := conf.InitConfig()
	fmt.Println(appConfig)
	fmt.Println(int(appConfig["logRotateSize"].(int64)))

	log := log2.NewLogger(appConfig)
	log.Finest("finest")
	log.Fine("fine")
	log.Debug("debug")
	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	log.Info("The time is now: %s", time.Now().Format("2006-01-02 15:04:05"))
	log.Warn("warn")
	log.Critical("critical")

	//var appConfig map[string]interface{}

	ln, err := net.Listen("tcp", "127.0.0.1:6010")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Critical("get client connection error: ", err)
		}
		go handleConnection(conn)
	}
}
