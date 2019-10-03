package taskrunner

import "log"

type Runner struct {
	Controller controlChan
	Error controlChan
	Data dataChan
	dataSize int
	longLived bool
	Dispatcher fn
	Executor fn
}

func NewRunner(size int, longLived bool, d fn, e fn) *Runner  {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		longLived:  longLived,
		dataSize:   size,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatch()  {
	defer func() {
		//这个参数为判断是否要回收资源
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()
	for ;; {
		select {
		case c := <- r.Controller :
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					runnerForError(r, err)
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}
			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					runnerForError(r, err)
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <- r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}
}

func (r *Runner) startAll()  {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}

func runnerForError(r *Runner, err error)  {
	log.Println(err)
	r.Error <- CLOSE
}