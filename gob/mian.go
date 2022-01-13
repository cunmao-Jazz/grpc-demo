package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func GobEncode(val interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(val); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func GobDecode(data []byte, value interface{}) error {
	reader := bytes.NewReader(data)
	decoder := gob.NewDecoder(reader)
	return decoder.Decode(value)
}

type TestStruct struct {
	name  string
	Value string
}

func main() {
	t1 := &TestStruct{"name", "value"}
	resp, err := GobEncode(t1)
	fmt.Println(resp, err)

	t2 := &TestStruct{}
	GobDecode(resp, t2)
	fmt.Println(t2, err)
}
