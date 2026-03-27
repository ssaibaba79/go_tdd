package main

import (
	"fmt"
	"net/http"
)

type Health struct {
	Url string
	Ok  bool
}

type Results map[string]Health

func HealthCheckWorker(in <-chan string, out chan<- Health, done chan<- bool, check func(string) bool) {

	for url := range in {
		fmt.Println("HealthCheckWorker processing :", url)
		ok := check(url)
		out <- Health{url, ok}
	}

	done <- true
}

func ResultCollector(in <-chan Health, out chan<- Results) {
	defer close(out)
	//fmt.Println("collecting started")
	results := make(map[string]Health)
	for health := range in {
		fmt.Printf("collecting for %v \n", health)
		results[health.Url] = health
	}
	//fmt.Printf("Writing results %v \n", results)
	out <- results
	//fmt.Printf("Written results %v \n", results)
}

func main() {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"http://bad.site.x",
	}

	input := make(chan string)
	health := make(chan Health)
	results := make(chan Results)
	done := make(chan bool)

	fmt.Println("Staring workers")
	for i := 0; i < len(websites); i++ {
		go HealthCheckWorker(input, health, done, func(url string) bool {
			r, err := http.Get(url)
			//fmt.Printf("check for %s returned %v, %v \n", url, r, err)
			if err == nil {
				//fmt.Printf("check for %s returned %v \n", url, r.StatusCode)
				return r.StatusCode == 200
			} else {
				//fmt.Printf("check for %s returned %v \n", url, err)
				return false
			}
		})
	}
	fmt.Println("Staring collector")
	go ResultCollector(health, results)

	// produce input
	for _, w := range websites {
		input <- w
	}

	// done sending input
	close(input)

	// check all is processed
	for i := 0; i < len(websites); i++ {
		<-done
	}
	close(health)

	// fetch results
	r := <-results
	fmt.Printf("Result %#v", r)

}
