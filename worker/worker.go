package worker


type Itask interface {
	Task(successOutput chan interface{}, errorOutpur chan error)
}

type Workers struct {

}



func (w Workers) DefineWorker(n int) []Workers{
	wout := make([]Workers,0)
	for {
		if n < 1 {
			break
		}
		wout = append(wout, Workers{})
		n--
	}
	return wout
}

func (w *Workers) TugasBersama(n int, task Itask) (<-chan interface{}, <-chan error) {
	outch := make(chan interface{})
	errch := make(chan error)
	ws := w.DefineWorker(n)
	for range ws {
		go task.Task(outch, errch)
	}
	return outch, errch
}