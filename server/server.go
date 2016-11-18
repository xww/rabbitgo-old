package main

import (
	"bufio"
	"fmt"
	"github.com/xww/rabbitgo/log2"
	"log"
	"net"
	"time"
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

	log := log2.NewLogger()
	log.Finest("Everything is created now (notice that I will not be printing to the file)")
	log.Fine("aaa")
	log.Debug("debug")
	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	log.Warn("warn")
	log.Critical("Time to close out!")

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
