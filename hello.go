package main

import "fmt"

const englishGreetHello = "Hello"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s", englishGreetHello, name)
}

func main() {

	fmt.Println(Hello("Sarvan"))
}
