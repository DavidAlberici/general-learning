package main

import "fmt"

func GetHello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(GetHello("user"))
}
