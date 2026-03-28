package main

import (
	"fmt"
	"time"
)

func main() {

	c := ticker("test", 500)
	for i := 0; i < 10; i++ {
		t := <-c
		fmt.Printf("Tick %d %s \n", i, t)
	}

}

type Tick struct {
	Name string
	Time time.Time
	Next time.Time
}

func (t Tick) String() string {
	FORMAT := "15:04:05.000"
	return fmt.Sprintf("[%s %s next at %s ]", t.Name, t.Time.Format(FORMAT), t.Next.Format(FORMAT))
}

func ticker(name string, tickIntervalInMilliSeconds int) chan Tick {
	c := make(chan Tick)

	go func() {
		for {
			now := time.Now()
			delay := time.Duration(tickIntervalInMilliSeconds) * time.Millisecond
			nextTick := now.Add(delay)
			tick := Tick{name, now, nextTick}
			c <- tick
			time.Sleep(delay)
		}
	}()

	return c

}
