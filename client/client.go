package main

import (
	"bufio"
	"fmt"
	"net"
	"bytes"
	"encoding/binary"
	"strings"
	"time"
	//"unsafe"
	//"crypto/sha1"
	"crypto/sha1"
	"github.com/zhuangsirui/binpacker"

)


type CmdProtocol struct {
	Type int32
	Version uint32
	Token uint32
	Length uint32
	MsgContent string
	CheckSum uint32
}

func testbytetrans()  {
	b := []byte{0x3, 0x32, 0x3, 0x32}
	fmt.Println(b)
	b_buf := bytes.NewBuffer(b)
	var x int32
	binary.Read(b_buf, binary.BigEndian, &x)
	fmt.Println(x)
	fmt.Println(strings.Repeat("-", 100))

	c := int16(1000)
	b_buf = bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, c)
	fmt.Println(b_buf.Len())
	binary.Write(b_buf, binary.BigEndian, c)
	fmt.Println(b_buf.Len())
	fmt.Println(b_buf.Bytes())

	d := "hello"
	d_byte := []byte(d)
	fmt.Println(d_byte)
	d_string := string(d_byte)
	fmt.Println(d_string)
}

func bb(){
	buf := new(bytes.Buffer)
	var data = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("%x", buf.Bytes())
	fmt.Println(buf.Bytes())
}

type T struct {
	A int16
	B int8
	//C []byte
}

func test1() {
	// Create a struct and write it.
	t := T{A: 99, B: 10}
	buf := &bytes.Buffer{}

	buf1 := []byte{100, 100}
	fmt.Println(buf1)

	buf.Write(buf1)

	//err := binary.Write(buf, binary.BigEndian, t)

	//if err != nil {
	//  panic(err)
	//}
	fmt.Println(buf)

	// Read into an empty struct.
	t = T{}
	err := binary.Read(buf, binary.BigEndian, &t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d %d", t.A, t.B)
}

type myStruct struct {
	ID   string
	Data string
}

func test2() {
	var bin_buf bytes.Buffer
	x := myStruct{"1", "Hello"}
	binary.Write(&bin_buf, binary.BigEndian, x)
	fmt.Println(bin_buf.Bytes())
	fmt.Printf("% x", bin_buf.Bytes())
	fmt.Printf("% x", sha1.Sum(bin_buf.Bytes()))
}
type TT struct {
	A uint16
	B string
	C []byte
}


func test3() {
	field1 := uint16(1)
	field2 := "Hello World"
	field3 := []byte("Hello World")
	buffer := new(bytes.Buffer)
	binpacker.NewPacker(buffer).
		PushUint16(field1).
		PushUint16(uint16(len(field2))).PushString(field2).
		PushUint16(uint16(len(field3))).PushBytes(field3)
	fmt.Println(buffer.Bytes())

	t := new(TT)
	unpacker := binpacker.NewUnpacker(buffer)
	unpacker.FetchUint16(&t.A).StringWithUint16Perfix(&t.B).BytesWithUint16Perfix(&t.C)
	fmt.Println(t)
	fmt.Println(t.A)
	fmt.Println(t.B)
	fmt.Println(t.C)
}

func main() {
	//test3()
	p := new(CmdProtocol)
	p.Type=1
	p.Version=2
	p.Token=123456789
	p.Length=7
	p.MsgContent="abcdefg"
	p.CheckSum=654321


	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, p.Type)
	binary.Write(b_buf, binary.BigEndian, p.Version)
	binary.Write(b_buf, binary.BigEndian, p.Length)
	binary.Write(b_buf, binary.BigEndian, p.Token)
	binary.Write(b_buf, binary.BigEndian, []byte(p.MsgContent))
	binary.Write(b_buf, binary.BigEndian, p.CheckSum)
	fmt.Println(b_buf.Bytes())

	buffer := new(bytes.Buffer)
	binpacker.NewPacker(buffer).PushInt32(p.Type).PushUint32(p.Version).
	PushUint32(p.Token).PushUint32(p.Length).PushString(p.MsgContent).
		PushUint32(p.CheckSum)
	fmt.Println(buffer.Bytes())

	cmdprotocol := new(CmdProtocol)
	unpacker := binpacker.NewUnpacker(buffer)
	unpacker.FetchInt32(&cmdprotocol.Type).FetchUint32(&cmdprotocol.Version).
	FetchUint32(&cmdprotocol.Token).FetchUint32(&cmdprotocol.Length).
		FetchString(uint64(cmdprotocol.Length),&cmdprotocol.MsgContent).
	FetchUint32(&cmdprotocol.CheckSum)
	fmt.Println(cmdprotocol)


	conn, err := net.Dial("tcp", "127.0.0.1:6010")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Fprintf(conn, b_buf.Bytes())
	recv_byt := make([]byte,1024)
	data, err := bufio.NewReader(conn).Read(recv_byt)
	fmt.Println(string(recv_byt))
	conn.Write(b_buf.Bytes())
	time.Sleep(time.Second * 10)
	conn.Write(b_buf.Bytes())

	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", data)
}
