package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)



func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	fmt.Print(buf.Bytes())
	return buf.Bytes(), nil
}

func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

type stu struct {
	Age int
	Name string

}

func main() {
	var s *stu = new(stu)
	s.Age = 11
	s.Name = "name"

	byteencode, _ := Encode(s)
	fmt.Printf("%s",byteencode)
}
