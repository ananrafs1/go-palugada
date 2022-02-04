package main

import (
	"fmt"
	"github.com/ananrafs1/go-palugada/worker"
	"sync"
	"time"
	"math/rand"
)


func main() {
	var sg sync.WaitGroup
	for i:= 0; i < 100; i++ {
		sg.Add(1)
		go func(i int) {
			defer sg.Done()
			outch, _ := (&worker.Workers{}).TugasBersama(5, &printer{})
			for output := range outch {
				print(fmt.Sprintf("%s %d", output.(string), i))
			}
		}(i)
	}
	sg.Wait()

	Workers := make(Workers,0)
}


type printer struct{}
func (p *printer) Task(successOutput chan interface{}, errorOutpur chan error, close func()){
	defer close()
	rand.Seed(time.Now().UnixNano())
	n := 2 + rand.Intn(5-2)
	time.Sleep(time.Duration(n)* time.Second)
	successOutput <- "TUGAS DISELESAIKAN "
}

func print(words string){
	fmt.Println(words)
}