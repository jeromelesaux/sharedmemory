package main

import (
	"fmt"

	"github.com/jeromelesaux/sharedmemory/model"
)

type Data struct {
	S string
}

func (d Data) String() string {
	return d.S
}

func main() {
	sm, err := model.NewInMemory(0)
	if err != nil {
		panic(err)
	}
	data := model.NewRing()
	data.Add(Data{S: "hello world"})
	b, err := data.ToBytes()
	if err != nil {
		panic(err)
	}
	if err := sm.Set(b); err != nil {
		panic(err)
	}

	fmt.Println(sm.String())

}
