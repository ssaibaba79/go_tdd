package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {

	o := fanIn(producer("Source1"), producer("Source2"), 500)
	for i := 0; i < 100; i++ {
		fmt.Println(<-o)
	}

}

func producer(source string) <-chan string {

	c := make(chan string)

	// producer
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s Event %d", source, i)
			time.Sleep(time.Duration(rand.Int64N(1000)) * time.Millisecond)
		}
	}()
	return c
}

// Fan-in source1 and source2 into output channel
// select construct
// timeout pattern
func fanIn(source1, source2 <-chan string, timeout int) <-chan string {
	o := make(chan string)

	go func() {
		for {
			select {
			case s := <-source1:
				o <- s
			case s := <-source2:
				o <- s
			case <-time.After(time.Duration(timeout) * time.Millisecond):
				fmt.Printf("Both sources timed out after %d milliseconds \n", timeout)
			}
		}
	}()
	return o
}
