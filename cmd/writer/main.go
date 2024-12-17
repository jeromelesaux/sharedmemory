package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/jeromelesaux/sharedmemory/model"
)

type Data struct {
	S string
}

func (d Data) String() string {
	return d.S
}

var (
	pid = flag.Int("ID of the shared memory", 0, "ID of the shared memory to read")
)

func main() {
	flag.Parse()
	sm, err := model.NewInMemory(*pid)
	if err != nil {
		panic(err)
	}
	data := model.NewRing()
	for i := 0; i < rand.IntN(100); i++ {
		data.Add(Data{S: "hello world"})
		data.Add(Data{S: "how are you"})
	}
	now := time.Now()
	b, err := data.ToBytes()
	if err != nil {
		panic(err)
	}
	if err := sm.Set(b); err != nil {
		panic(err)
	}
	fmt.Printf("write in [%d] milliseconds\n", time.Now().UnixMilli()-now.UnixMilli())
	fmt.Println(sm.String())
}
