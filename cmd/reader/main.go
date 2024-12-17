package main

import (
	"fmt"

	"github.com/jeromelesaux/sharedmemory/model"
)

type RData struct {
	S string
}

func (r RData) String() string {
	return r.S
}

func main() {
	sm, err := model.NewInMemory(65545)
	if err != nil {
		panic(err)
	}

	b, err := sm.Get()
	if err != nil {
		panic(err)
	}

	data := model.NewRing()
	if err := data.FromBytes(b); err != nil {
		panic(err)
	}

	fmt.Println(data.String())
}
