package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

var totalVotes int
var voteMutex sync.Mutex

func generateVotes(count int) {
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < count; i++ {
        randomVotes := rand.Intn(10) + 1
        
        voteMutex.Lock()
        totalVotes += randomVotes
        voteMutex.Unlock()
        
        time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
    }
}

func main() {
    var wg sync.WaitGroup
    numGenerators := 10
    wg.Add(numGenerators)

    for i := 0; i < numGenerators; i++ {
        go func() {
            generateVotes(10)
            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Printf("Итоговое количество голосов: %d\n", totalVotes)
}