package main

import (
    "fmt"
    "sync"
)

type Pool struct {
    workQueue chan string
    wg        sync.WaitGroup
}

func CreatePool(size int) *Pool {
    p := &Pool{
        workQueue: make(chan string, size),
    }
    for i := 0; i < size; i++ {
        go p.work()
    }
    return p
}

func (p *Pool) work() {
    for task := range p.workQueue {
        fmt.Println("Обработка:", task)
        p.wg.Done()
    }
}

func (p *Pool) SubmitJob(job string) {
    p.wg.Add(1)
    p.workQueue <- job
}

func (p *Pool) Finish() {
    close(p.workQueue)
    p.wg.Wait()
}

func main() {
    pool := CreatePool(3)
    for i := 0; i < 10; i++ {
        job := fmt.Sprintf("Задача-%d", i+1)
        pool.SubmitJob(job)
    }
    pool.Finish()
}