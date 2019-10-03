package taskrunner

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestRunner(t *testing.T)  {
	d := func(dc dataChan) error {
		for i:= 0; i < 30; i++ {
			dc <- i;
			log.Println("Dispatcher send : ", i)
		}

		return nil
	}

	e := func(dc dataChan) error {
		forloop:
			for ; ;  {
				select {
				case d := <- dc:
					log.Println("Executor received", d)
				default:
					break forloop
				}
			}
			return errors.New("Executor end")
		}
	runner := NewRunner(30, false, d, e)
	go runner.startAll()
	time.Sleep(3 * time.Second)
}

func TestMap(t *testing.T)  {
	m := &sync.Map{}
	go func(p *sync.Map) {
		p.Store("hello", "world")
	}(m)
	time.Sleep(1 * time.Second)
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, " --> ", value)
		return true
	})
}