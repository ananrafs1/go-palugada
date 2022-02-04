package worker

import(
	"sync"
	"context"
	"fmt"
)
type Itask interface {
	Task(successOutput chan interface{}, errorOutpur chan error, close func())
}

type Worker struct {
	Task Itask
	OutputChan chan interface{}
	ErrorChan chan error
}



func (w Worker) DefineWorker(n int) []Worker{
	wout := make([]Worker,0)
	for {
		if n < 1 {
			break
		}
		wout = append(wout, Worker{})
		n--
	}
	return wout
}

func (w *Worker) TugasBersama(n int, task Itask) ( chan interface{},  chan error) {
	outch := make(chan interface{}, n)
	errch := make(chan error, n)
	ws := w.DefineWorker(n)
	var syncG sync.WaitGroup
	for range ws {
		syncG.Add(1)
		go task.Task(outch, errch, func() { syncG.Done() })
	}
	syncG.Wait()
	close(outch)
	close(errch)
	return outch, errch
}

func (w *Worker) Do (ctx context.Context, close func())  {
	go w.Task.Task(w.OutputChan, w.ErrorChan, close)
}

type Workers []Worker

func (w *Workers) DoAsync (ctx context.Context) {
	// var syncG sync.WaitGroup
	for ws := range *w {
		fmt.Println(ws)
		// syncG.Add(1)
		// (&ws).Do(ctx, func() {syncG.Done()})
	}
	// syncG.Wait()
}