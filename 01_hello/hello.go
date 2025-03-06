package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch {
	case lang == french:
		prefix = frenchHelloPrefix
	case lang == spanish:
		prefix = spanishHelloPrefix

	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("dude", "Spanish"))
}
