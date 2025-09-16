package naloge

import (
	"fmt"
	"strings"
)

func NalogeIzPrintanja() {
	err := fmt.Errorf("oopsie-daisy")
	formatted := fmt.Sprintf("integers: %d, floats: %f, strings: %s, quoted: %q, verbose: %v", 10, 3.14, "Hello World!", "quoted text", err)
	fmt.Println(formatted)

	// 1. Print a Greeting
	// Task: Write a program that prints: Hello, <your name>! Welcome to Go.

	name := "Nejc"
	fmt.Println("Hello, " + name + "! Welcome to Go.")

	// 2. Concentrate Strings
	// Task: Combine first and last name into a full name and print

	lastname := "Cencljotka"

	fullname := fmt.Sprintf("Ime in Priimek: %s %s", name, lastname)
	fmt.Println(fullname)

	// 3. String Length
	// Task: Ask for a string input and print how many characters it has.
	// inputString := "some string foobar"

	inputString := "some string foobar"
	fmt.Println("inputString has:", len(inputString), "characters")

	// 4. Change Case
	// Task: Convert a string to uppercase and lowercase and print both:
	// lowercaseString := "im a lowercase"
	// uppercaseString := "IM AN UPPERCASE"

	lowercaseString := "im a lowercase"

	uppercaseString := "IM AN UPPERCASE"

	// lowerCase := strings.ToUpper(lowercaseString)
	// upperCase := strings.ToLower(uppercaseString)
	// fmt.Println(lowecase)
	// fmt.Println(upperCase)

	fmt.Println(strings.ToUpper(lowercaseString))
	fmt.Println(strings.ToLower(uppercaseString))

	// 5. Replace Words
	// Task: Replace the word "boring" with "awesome" in a string like:
	// Go is boring. -> Go is awesome.

	opinion := "Go is boring"

	if opinion == "Go is boring" {
		opinion = "Go is awesome"
	}

	fmt.Println(opinion)

	originalString := "Go is boring"
	replacedString := strings.ReplaceAll(originalString, "boring", "awesome")

	fmt.Println(replacedString)

	// 6. Format Output
	// Task: Format and print a sentence like:
	// Name: Aljoša | Age: 23 | Country: Serbia
	// name := "Aljoša"
	// age := 23
	// country := "Serbia"

	name = "Aljoša"
	age := 23
	country := "Serbia"

	fmt.Printf("Name: %s, | Age: %d, | Country: %s\n", name, age, country)

	// 7. Print Special Characters
	// Instruction: Print a sentence with quotation marks and a line break

	fmt.Printf("%q, %s, %s\n", "quote", "se uporablja k quotaš", "\\")

	// 8: Combine Multiple Variables into a Sentence
	// Instruction: Use several string variables to make and print a sentence.
	// a := "The"
	// b := "sky"
	// c := "is"
	// d := "blue"

	a := "The"
	b := "sky"
	c := "is"
	d := "blue"

	fmt.Printf("%s %s %s %s\n", a, b, c, d)

}
