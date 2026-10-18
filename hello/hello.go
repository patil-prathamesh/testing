package hello

import "fmt"

const englishPrefixHello = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "spanish" {
		return "Hola, " + name
	}
	return englishPrefixHello + name
}

func main() {
	fmt.Println(Hello("Hello", "english"))
}
