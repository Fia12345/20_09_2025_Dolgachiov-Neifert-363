package main
import (
	"fmt"
	"time"
	"sync"
)

type Cinema struct {
	seats []bool
	mutex sync.Mutex
}

func (c.Cinema) bookSeat(seat int, user string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.seats[seat] {
		return false
	}

	c.seats[seat] = true
	fmt.Printf("%s Забронированно метс %d\n", user, seat+1)
	return true
}

func main() {
	cinema := &cinema{seats:make([]bool, 38)}
	var wg sync.WaitGroup

	for i := 1; i<= 50; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)

			seat := userID % 38
			if !cinema.bookSeat(seat, fmt.Sprintf("User%d", userID)) {
				fmt.Printf("User%d не смог забонировать место %d\n", userID, seat+1)
			}
		}(i)
	}

	wg.Wait()

	booked := 0
	for i, taken := range cinema.seats {
		if taken {
			booked++
			fmt.Printf("место %d забранировано\n", i+1)
		}
	}
	fmt.Printf("\nЗабронировано мест:%d/%d\n", booked,len(cinema.seats))
}