package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gromnsk/slidingWindow"
)

func main() {
	lim := slidingWindow.NewLimiter(time.Second, 25*time.Millisecond, 100)
	counter := 0
	now := time.Now()
	for i := 0; i < 1000; i++ {
		r := rand.Intn(15)
		time.Sleep(time.Duration(r) * time.Millisecond)
		if lim.Allow() {
			lim.Take()
			counter++
		}
	}
	fmt.Println(time.Now().Sub(now))
	fmt.Println(counter)
}
