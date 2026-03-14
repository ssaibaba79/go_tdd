package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
  CountDown(os.Stdout, DefaultSecondsSleeper{})
}

func CountDown(w io.Writer, sleeper Sleeper) {

	for i:=3;i>0;i--{
		fmt.Fprintf(w, "%d\n", i)
	}
	sleeper.Sleep()

	fmt.Fprintf(w, "Go!")
}

type Sleeper interface{
	Sleep()
}

type DefaultSecondsSleeper struct {}

func (d DefaultSecondsSleeper ) Sleep(){
	time.Sleep(1 * time.Second)
}




