package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishGreetPrefix = "Hello"
const spanishGreetPrefix = "Hola"
const frenchGreetPrefix = "Bonjour"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishGreetPrefix
	case french:
		prefix = frenchGreetPrefix
	default:
			prefix = englishGreetPrefix
	}
	return
}

func main() {
}
