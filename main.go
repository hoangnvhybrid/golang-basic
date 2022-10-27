package main

import "fmt"

func main() {
	msg := sayHello("Jon")
	fmt.Println(msg)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
