package main

import (
	"fmt"
	"github.com/ananrafs1/go-palugada/worker"
)


func main() {
	o, _ := Workers{}.TugasBersama(5, printer{})
	for range o {
		print(o)
	}
}


type printer struct{}
func (p *printer) Task(successOutput chan interface{}, errorOutpur chan error){
	successOutput <- "TUGAS DISELESAIKAN "
}

func print(words string){
	fmt.Println(words)
}