package main

import (
	"fmt"
	"github.com/ananrafs1/go-palugada/worker"
	// "sync"
	"time"
	"math/rand"
)


func main() {
	
	Tasks := []worker.Itask{ printer{}, printer{}, printer{}, printer{} }
	Workers := worker.Worker{
			Task : Tasks,
		}
	Workers.Do()
	ret, _ := Workers.Listen()
	for ls := range ret {
		fmt.Println(ls.(string))
	}
}


type printer struct{
	worker.Itask
}
func (p printer) Task() (interface{},  error) {
	rand.Seed(time.Now().UnixNano())
	n := 2 + rand.Intn(5-2)
	time.Sleep(time.Duration(n)* time.Second)
	return "TUGAS DISELESAIKAN ", nil
}
