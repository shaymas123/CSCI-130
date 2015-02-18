// STEP 1 - GOAL:
// to be able to add *any* string to the end of what is printed

package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

type MyPrinterType func(string)

func Greet(person Contact, myFunc MyPrinterType) {
	greeting, name := CreateMessage(person.name, person.greeting, "Good day")
	myFunc(greeting)
	myFunc(name)
}

func CreateMessage(name string, greeting ...string) (myGreeting string, myName string) {
	myGreeting = greeting[1] + " " + name
	myName = "\nHey, " + name + "\n"
	return
}

func myPrint(s string) {
	fmt.Print(s)
}

func myPrintln(s string) {
	fmt.Println(s)
}

func myPrintCustom(s string, custom string) {
	fmt.Println(s + custom)
}

func myPrintFunction(custom string) MyPrinterType {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func main() {

	var t = Contact{}
	t.greeting = "Hello"
	t.name = "Shay"

	Greet(t, myPrintFunction("!!!"))

	switch t.name {
	case "bob":
		fmt.Println("wrong")
	case "Jim":
		fmt.Println("wrong again")
	case "Shay":
		fmt.Println("Right!")
	default:
		fmt.Println("Have you no friends?")
	}

	if t.greeting == "Hello" {
		fmt.Println("You got it right")
		fmt.Println("Good job")
	} else {
		fmt.Println("You got it wrong")
	}

	if t.greeting == "Goodbye" {
		fmt.Println("You got it right")
		fmt.Println("Good job")
	} else {
		fmt.Println("You got it wrong")
	}

}
