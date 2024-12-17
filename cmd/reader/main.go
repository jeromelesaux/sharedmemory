package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jeromelesaux/sharedmemory/model"
)

type RData struct {
	S string
}

func (r RData) String() string {
	return r.S
}

var (
	pid = flag.Int("id", 0, "ID of the shared memory to read")
)

func main() {
	flag.Parse()
	sm, err := model.NewInMemory(*pid)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	b, err := sm.Get()
	if err != nil {
		panic(err)
	}

	data := model.NewRing()
	if err := data.FromBytes(b); err != nil {
		panic(err)
	}
	fmt.Printf("read in [%d]\n", time.Now().UnixMilli()-now.UnixMilli())

	// fmt.Println(data.String())
}
