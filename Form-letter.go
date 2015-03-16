package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
Dear {{.Name}},
{{if .Attended}}
I'm glad you could make it to the wedding!!!.{{else}}
It is a shame you couldn't make it to the wedding.{{end}}
{{with .AwesomeGift}}Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Shaymas
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, AwesomeGift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Rachel Souza", "Fork", true},
		{"Shaymas Irving", "Cup", false},
		{"Brandon Davis", "Spoon", false},
	}

	// STEP 1 & STEP 2
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	//STEP 3
	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}