package main

import "fmt"

func GetHello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(GetHello("user"))
}
