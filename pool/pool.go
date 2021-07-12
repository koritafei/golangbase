package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 实现一个通用的协程池

// 数据结构
type Task struct {
	Data    []interface{}
	Handler func(v ...interface{})
}

var ErrPoolCap = errors.New("capacity num is not error")
var ErrPoolClosed = errors.New("pool has been closed")

type Pool struct {
	// 容量
	capacity uint64
	// 任务队列
	taskC chan *Task
	// 正在运行协程数
	runningTaskCount uint64
	// 协程池状态
	state int64

	// panic recover Handler
	panicRecoverHandler func(v ...interface{})

	// 互斥锁
	mutex sync.Mutex
}

const (
	STOPED  = 0
	RUNNING = 1
)

// create pool
func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrPoolClosed
	}

	return &Pool{
		capacity:         capacity,
		state:            RUNNING,
		runningTaskCount: 0,
		taskC:            make(chan *Task, capacity),
	}, nil
}

// inc runningTaskCount
func (p *Pool) incRunningTaskCount() {
	atomic.AddUint64(&p.runningTaskCount, 1)
}

// dec runningTaskCount
func (p *Pool) decRunningTaskCount() {
	atomic.AddUint64(&p.runningTaskCount, ^uint64(0))
}

// get runningTaskCount
func (p *Pool) getRunningTaskCount() uint64 {
	return atomic.LoadUint64(&p.runningTaskCount)
}

// get capacity
func (p *Pool) GetCapacity() uint64 {
	return p.capacity
}

// get statement
func (p *Pool) getStatement() int64 {
	return atomic.LoadInt64(&p.state)
}

// set statement
func (p *Pool) setStatement(state int64) {
	atomic.StoreInt64(&p.state, state)
}

// Put task
func (p *Pool) Put(task *Task) error {
	fmt.Println("Put Task To Pool")
	if p.getStatement() == STOPED {
		return ErrPoolClosed
	}

	p.mutex.Lock()
	if p.GetCapacity() >= p.getRunningTaskCount() {
		p.run()
	}
	p.mutex.Unlock()

	p.mutex.Lock()
	if p.getStatement() == RUNNING {
		// safe push task
		fmt.Println("Put TaskC")
		p.taskC <- task
	}
	p.mutex.Unlock()

	return nil
}

// start pool
func (p *Pool) run() {
	fmt.Println("Run Task")
	p.incRunningTaskCount()

	go func() {

		defer func() {
			p.decRunningTaskCount()
			if r := recover(); r != nil { // recover panic
				if p.panicRecoverHandler != nil {
					p.panicRecoverHandler(r)
				} else {
					fmt.Println("panic recover handler not set")
				}
			}
		}()

		for {
			select {
			case task, ok := <-p.taskC:
				if !ok {
					// closed
					return
				}
				fmt.Println("Run Task Handler")
				task.Handler(task.Data...)
			}
		}

	}()

}

// close
func (p *Pool) Close() {
	if p.state == STOPED {
		return
	}
	p.setStatement(STOPED)

	p.close()
}

func (p *Pool) close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for len(p.taskC) > 0 {
		// 阻塞，等待任务全部处理
	}

	close(p.taskC)
}

func main() {

	p, err := NewPool(20)

	if err != nil {
		fmt.Println("Error Initial " + err.Error())
		panic(err)
	}

	for i := 0; i < 50; i++ {
		p.Put(&Task{
			Data: []interface{}{i},
			Handler: func(v ...interface{}) {
				fmt.Println(v)
			},
		})
	}

	time.Sleep(time.Second * 10)

}
