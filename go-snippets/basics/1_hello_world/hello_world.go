package main

import (
	"fmt"
	"os"
)

const (
	englishPrefix = "Hello, "
	spanishPrefix = "¡Hola, "
	frenchPrefix  = "Salut, "

	spanish = "ES"
	french  = "FR"
)

func GetHello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	// argsWithProgramPath := os.Args
	args := os.Args[1:]

	name := ""
	language := ""

	if len(args) == 2 {
		name = args[0]
		language = args[1]
	}

	fmt.Println(GetHello(name, language))
}
