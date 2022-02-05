package worker

import(
	"sync"
	// "context"
	// "fmt"
)
type Itask interface {
	Task() (ret interface{}, err error)
	Write() (ret interface{}, err error)
}

type Worker struct {
	Task []Itask
	outputChan chan interface{}
	errorChan chan error
}


func (w *Worker) Do() {
	if len(w.Task) < 1 {
		return
	}
	w.outputChan = make(chan interface{}, len(w.Task))
	w.errorChan = make(chan error, len(w.Task))
	var wg sync.WaitGroup
	for _,t := range w.Task {
		wg.Add(1)
		go func(){
			defer wg.Done()
			ret, err := t.Task()
			if err != nil {
				w.errorChan <- err
				return
			}
			w.outputChan <- ret
		}()
	}
	wg.Wait()
	close(w.outputChan)
	close(w.errorChan)
}


func (w *Worker) Listen() (<-chan interface{}, <-chan error) {
	return w.outputChan, w.errorChan
}