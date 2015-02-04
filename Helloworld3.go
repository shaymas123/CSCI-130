package main

import "fmt"

type salutation struct {
	name     string
	greeting string
}

func greet(salutation salutation) {
	fmt.Println(salutation.name)
	fmt.Println(salutation.greeting)
}
func main() {
	var s = salutation{"Bob", "Hello"}
	greet(s)
}
