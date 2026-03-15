package main

import (
	"fmt"
	"log"
	"net/http"
)

type chanInt chan int

func (ch chanInt) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"next :%d", <-ch)
}


func counter (ch chan<-int) {
  for i:=0;;i++{
		ch <- i
	}
}



func main(){
  var nextId chanInt = make(chan int)
	go counter(nextId)
	
	http.HandleFunc("/", nextId.handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

