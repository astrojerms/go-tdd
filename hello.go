package main

import "fmt"

// How do you test this? It is good to separate your "domain" code from the outside world (side-effects).
// The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	Spanish            = "Spanish"
	English            = "English"
	French             = "French"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greatingPrefix(language) + name + "."
}

func greatingPrefix(language string) (prefix string) {
	switch language {
	case Spanish:
		prefix = spanishHelloPrefix
	case French:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("human-person.", "English"))
}
