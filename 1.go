package main

import (
	"fmt"
	"sync/atomic"
	"sync"
)

var counter int64

func main() {
	
	var wg sync.WaitGroup
	gorut := 67

	for i:= 0; i < gorut; i++{
		wg.Add(1)
		go func(){
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	} 
	
	wg.Wait()
	
	fmt.Println("Ну вроде ответ:", atomic.LoadInt64(&counter))
}