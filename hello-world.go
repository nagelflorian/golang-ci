package main

import "fmt"

// Greeting : returns a formatted greeting string
func Greeting(s string) string {
	return fmt.Sprintf("Hi, %s", s)
}

func main() {
	greeting := Greeting("Felix")
	fmt.Println(greeting)
}
