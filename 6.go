package main

import (
    "fmt"
    "sync"
)

var logBuffer []string
var lock sync.Mutex

func writeLog(msg string) {
    lock.Lock()
    defer lock.Unlock()
    logBuffer = append(logBuffer, msg)
    fmt.Println(msg)
}

func main() {
    var waiter sync.WaitGroup
    for i := 0; i < 10; i++ {
        waiter.Add(1)
        go func(i int) {
            writeLog(fmt.Sprint("Логирование ", i))
            waiter.Done()
        }(i)
    }
    waiter.Wait()
}