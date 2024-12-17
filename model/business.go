package model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Ring struct {
	L    LinkedList
	Name string
}

func NewRing() Ring {
	return Ring{
		L:    NewLinkedList(),
		Name: "RingStructure",
	}
}

func (r Ring) ToBytes() ([]byte, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(&r)
	return buf.Bytes(), err
}

func (r *Ring) FromBytes(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(r)
}

func (r *Ring) Add(data any) {
	r.L.Add(data)
}

func (r *Ring) String() string {
	var str string
	str += fmt.Sprintf("Name: [%s], Len: [%d]\n", r.Name, r.L.Len)
	e := r.L.Root.Next
	for i := 0; i < r.L.Len; i++ {
		val, ok := e.Value.(string)
		if !ok {
			str += fmt.Sprintf("Node: [%d], Value: [%s]\n", i, e.Value)
		} else {
			str += fmt.Sprintf("Node: [%d], Value: [%s]\n", i, val)
		}
		e = e.Next
	}
	return str
}
