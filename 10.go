package main
import (
	"errors"
	"fmt"
	"sync"
	"time"
)


type ErrorCollector struct {
	errors []error
	mutex sync.Mutex
}

func (ec *ErrorCollector) Add(err error) {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	ec.errors = append(ec.errors, err)
}

func (ec *ErrorCollector) GetAll() []error {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	return ec.errors
}

func generateData(nums []int) <- chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()
	return out
}

func multiplyByTwo(in <- chan int,ec *ErrorCollector) <- chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			if num == 3 {
				ec.Add(errors.New("Ошибка умножения для числа 3"))
				continue
			}
			result := num * 2
			out <- result
		}
	}()
	return out
}

func addTen(in <- chan int, ec*ErrorCollector) <- chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			if num > 15 {
				ec.Add(errors.New("Ошибка сложения для числа > 15"))
				continue
			}
			result := num + 10
			out <- result
		}
	}()
	return out
}


func main() {
	data := []int{1,2,3,4,5,6,7,8,9,10}
	errorCollector := &ErrorCollector{}

	stage1 := generateData(data)
	stage2 := multiplyByTwo(stage1, errorCollector)
	stage3 := addTen(stage2, errorCollector)

	var results []int
	for result := range stage3 {
		results = append(results, result)
		fmt.Printf("Результат: %d\n", result)
		time.Sleep(100 * time.Millisecond)
	}

	errors := errorCollector.GetAll()
	if len(errors) >0 {
		fmt.Printf("\n Собрано ошибок: %d\n", len(errors))
		for i, err := range errors {
			fmt.Printf("ошибка %d: %v \n",i+1,err)
		}
	}else {
		fmt.Println("\nОшибок не обнаружено")
	}
	fmt.Printf("Успешно обработанных  результатов: %d\n", len(results))
}
