package main

import (
	"github.com/golang/protobuf/proto"

	"fmt"
	"os"

	"github.com/xww/rabbitgo/protobuf/demo"
)

func main() {
	//注意每条信息后面的,号
	msg_test := &Msg{
		MsgType: proto.Int32(1),
		MsgInfo: proto.String("I am hahaya."),
		MsgFrom: proto.String("127.0.0.1"),
	}
	//将数据序列化到字符串中(写操作)
	in_data, err := proto.Marshal(msg_test)
	if err != nil {
		fmt.Println("Marshaling error: ", err)
		os.Exit(1)
	}
	//将数据从字符串中反序列化出来(读操作)
	msg_encoding := &demo.Msg{}
	err = proto.Unmarshal(in_data, msg_encoding)
	if err != nil {
		fmt.Println("Unmarshaling error: ", err)
		os.Exit(1)
	}
	fmt.Printf("msg type: %d\n", msg_encoding.GetMsgType())
	fmt.Printf("msg info: %s\n", msg_encoding.GetMsgInfo())
	fmt.Printf("msg from: %s\n", msg_encoding.GetMsgFrom())
}
