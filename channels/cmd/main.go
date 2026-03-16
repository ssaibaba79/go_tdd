package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func getCounts() {

	for {
		 log.Printf("Getting next int")
		r, _ := http.Get("http://localhost:8080/")
		body, err := io.ReadAll(r.Body)
	   if err != nil {
	      log.Fatalln(err)
	   }
	   sb := string(body)
	   log.Printf("%s", sb)

		  time.Sleep(time.Second)
		}
}

func main() {


 input := make(chan int)
 output := make(chan int)
 
 // 1 worker
 go func(input <-chan int, output chan<- int) {
			fmt.Println("in Goroutine")
			i := <-input
			output <- (i * 2)
	 }(input, output)

	// produce to input channel
 	input <- 1
	// receive from output channel 
	fmt.Println("result :", <-output )

  // example
	 getCounts() 
 } 
