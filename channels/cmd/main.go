package main

import (
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
 getCounts() 
}