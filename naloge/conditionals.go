package naloge

import (
	"fmt"
	"strconv"
)

func NalogeIzConditionals() {
	// 1. Check if the boolean variable x is true.
	// If it is, print "x is true".
	// x := true

	x := true
	if x {
		fmt.Println("1. x je true")
	} else {
		fmt.Println("1. x je false")
	}

	// 2. Determine whether the given number is even or odd.
	// Print "Even" or "Odd" accordingly.
	// number := 7

	number := 7

	if number % 2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// 3. Print the appropriate letter grade based on the score.
	// Use: A for 90+, B for 80-89, C for 70-79, F otherwise.
	// score := 85

	score := 85

	if score >= 90 {
		fmt.Println("A")
	} else if score >80 && score <89 {
		fmt.Println("B")
	} else if score >70 && score <79 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}


	// 4. Check whether the number is positive, negative, or zero.
	// Print the result.
	// n := -3

	n := -3

	if n < 0 {
		fmt.Println("negatif")
	} else if n == 0 {
		fmt.Println("zzzzzeeroooooo")
	} else if n > 0 {
		fmt.Println("Pozitiva eeoooo")
	}


	// 5. Compare two strings and print whether they are equal or different.
	// a := "hello"
	// b := "world"

	a := "hello"
	b := "world"

	if a == b {
		fmt.Println("egual")
	} else if a != b {
		fmt.Println("different")
	}


	// 6. Use a switch statement to print the day of the week based on a number (1-5).
	// Print "Weekend" for any other value.
	// day := 3

	day := 3

	if day == 1 {
		fmt.Println("Pondelk")
	} else if day == 2 {
		fmt.Println("tork")
	} else if day == 3 {
		fmt.Println("sredah")
	} else if day == 4 {
		fmt.Println("Četrvrtak")
	} else if day == 5 {
		fmt.Println("petekzadetek")
	} else {
		fmt.Println("vikendica")
	}

	switch day {
	case 1: // if day == 1
		fmt.Println("mon")
	case 2: // if day == 2
		fmt.Println("tue")
	case 3: // if day == 3
		fmt.Println("wed")
	case 4: // if day == 4
		fmt.Println("thu")
	case 5: // if day == 5
		fmt.Println("fri")
	default: // else
		fmt.Println("weekend")
	}

	// 7. Check whether the number is divisible by both 3 and 5.
	// num := 15

	num := 15

	if num % 3.0 == 0 && num % 5.0 == 0 {
		fmt.Println("It is divisible")
	} else {
		fmt.Println("it is NOT divisible")
	}

	// 8. Determine if a given year is a leap year.
	// A leap year is divisible by 4 but not by 100, unless also divisible by 400.
	// year := 2024

	year := 2024

	if year % 4 == 0 && year % 100 == 0 && year % 400 == 0 {
		fmt.Println("it is a leap year")
	} else if year % 4 == 0 && year % 100 != 0 {
		fmt.Println("it is a leap year")
	} else {
		fmt.Println("it is NOT a leap year")
	}

	// 9. Check if the given character is a vowel (a, e, i, o, u) or consonant.
	// ch := 'e'

	ch := "e"

	switch ch {
		case "a": 
			fmt.Println("True")
		case "e": 
			fmt.Println("True")
		case "i": 
			fmt.Println("True")
		case "o": 
			fmt.Println("True")
		case "u": 
			fmt.Println("True")
		default: 
			fmt.Println("false")
		}

	// 10. Check if a person is eligible to vote (age 18 or older).
	// age := 16

		age := 16

		if age >= 18 {
			fmt.Println("Eligible")
		} else {
			fmt.Println("NOT Eligible")
		}

	// 11. Use a switch to respond to a traffic light color.
	// For example: "Stop" for red, "Go" for green, "Wait" for yellow.
	// color := "green"

	color := "yellow"

	if color == "green" {
		fmt.Println("Go")
	} else if color == "red"{
		fmt.Println("Stop")
	} else if color == "yellow"{
		fmt.Println("Wait")
	}

	// 12. Compare three integers and print the largest one.
	// x := 3
	// y := 7
	// z := 5

	xx := 3
	y := 7
	z := 5

	rezultat := 0

	if xx > rezultat {
		rezultat = xx
	}
	if y > rezultat {
		rezultat = y
	}
	if z > rezultat {
		rezultat = z
	}
	fmt.Println(rezultat)

	list := []int{xx, y, z}
	res := list[0]
	for _, i := range list {
		if i > res {
			res = i
		}
	}

	// 13. Check if a number is within the range 10 to 20 (inclusive).
	// val := 15

	val := 15

	if val >= 10 && val <= 20 {
		fmt.Println("Inclusive")
	} else {
		fmt.Println("NON INCLUSIVE!!!")
	}

	// 14. Classify the temperature:
	// below 0: "Freezing", 0-20: "Cold", 21-30: "Warm", above 30: "Hot"
	// temperature := 30

	temperature := 30

	if temperature < 0 {
		fmt.Println("Freezing")
	} else if temperature > 0 && temperature < 20 {
		fmt.Println("Cold")
	} else if temperature > 21 && temperature < 30 {
		fmt.Println("Warm")
	} else {
		fmt.Println("Hot")
	}


	// 15. Use a switch statement to describe a grade:
	// A = "Excellent", B = "Good", C = "Average", F = "Fail"
	// grade := "B"

	grade := "A"

	switch grade {
		case "A": 
			fmt.Println("Excelent")
		case "B": 
			fmt.Println("Good")
		case "C": 
			fmt.Println("Average")
		case "F": 
			fmt.Println("Fail")
		}

	// 16. Check whether a given string is empty or not.
	// text := ""

		text := ""

		if text == "" {
			fmt.Println("Empty")
		} else {
			fmt.Println("Not empty")
		}

	// 17. Check if a number is divisible by 2, 5, both, or neither.
	// n := 10

		nuu := 10

		if nuu % 2 == 0 && nuu % 5 == 0 {
			fmt.Println("Both")
		}
		if nuu % 2 == 0 && nuu % 5 != 0 {
			fmt.Println("By 2")
		}
		if nuu % 2 != 0 && nuu % 5 == 0 {
			fmt.Println("By 5")
		}
		if nuu % 2 != 0 && nuu % 5 != 0 {
			fmt.Println("Neither")
		}

	// 18. Use a switch to print an emoji based on mood:
	// "happy" = 😀, "sad" = 😢, "angry" = 😡, default = 😐
	// mood := "happy"

		mood := "angry"

		switch mood {
			case "happy": 
				fmt.Println("😀")
			case "sad": 
				fmt.Println("😢")
			case "angry": 
				fmt.Println("😡")
			default: 
				fmt.Println("😐")
			}

	// 19. Check if a password is long enough (8 characters or more).
	// password := "secret"

		password := "secretttt"
		passs := len(password)

		if passs < 8 {
			fmt.Println("Not long enough!")
		} else {
			fmt.Println("Ok! Welcome")
		}

	// 20. Use if-else to determine if the given number is a multiple of 10.
	// number := 30

		nuu = 30

		if nuu % 10 == 0 {
			fmt.Println("It is!")
		} else {
			fmt.Println("It is NOT!")
		}

	// 21. Check if a number is odd AND greater than 10.
	// number := 13

		nuu = 13

		if nuu % 2 != 0 && nuu > 10 {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}

	// 22. Use a switch to simulate a basic calculator with operators "+", "-", "*", "/".
	// operator := "+"
	// a := 5
	// b := 3

		operator := "*"
		c := 5
		d := 3

	switch operator {
		case "+": 
			fmt.Println(c+d)
		case "-": 
			fmt.Println(c-d)
		case "*": 
			fmt.Println(c*d)
		case "/": 
			fmt.Println(c/d)
		}

	// 23. Determine if a number is a single-digit, two-digit, or larger.
	// number := 42

		nuu = 42

		strcon := strconv.Itoa(nuu)
		lengi := len(strcon)
		
		if lengi == 1 {
			fmt.Println("One-digit")
		} else if lengi == 2 {
			fmt.Println("Two-digit")
		} else {
			fmt.Println("Larger")
		}

	// 24. Print a special message if a user's name matches a secret name.
	// name := "Harry"

		name := "Harry"

		if name == "Harry" {
			fmt.Println("Happy Birthday")
		} else {
			fmt.Println("Nah man gtfo")
		}


	// 26. Check if two numbers are both positive.
	// a := 5
	// b := 12

		aa := 5
		bb := 12

		if aa > 0 && bb > 0 {
			fmt.Println("They are!")
		} else {
			fmt.Println("oh no! they are not. Sadly.")
		}

	// 27. Use a switch to determine the season based on the month number (1-12).
	// month := 4

		month := 5

		switch month {
			case 1: 
				fmt.Println("Jan")
			case 2: 
				fmt.Println("Feb")
			case 3: 
				fmt.Println("Mar")
			case 4: 
				fmt.Println("Apr")
			case 5: 
				fmt.Println("May")
			case 6: 
				fmt.Println("Jun")
			case 7: 
				fmt.Println("Jul")
			case 8: 
				fmt.Println("Aug")
			case 9: 
				fmt.Println("Sep")
			case 10: 
				fmt.Println("Oct")
			case 11: 
				fmt.Println("Nov")
			case 12: 
				fmt.Println("Dec")
			}

	// 28. Check if a given year is in the 20th or 21st century.
	// year := 1999

		year = 1999

		if year < 2000 {
			fmt.Println("20th century")
		} else if year < 2100 {
			fmt.Println("21th Century fox")
		}

	// 29. Check if a temperature is suitable for swimming (between 22 and 30°C).
	// temperature := 25

		temperature = 25

		if temperature > 22 && temperature < 30 {
			fmt.Println("It is a fun day fo a swim huh?")
		} else {
			fmt.Println("stay the fuck out man!")
		}

	// 30. Use a switch to simulate a mini menu system:
	// "1" = "Start Game", "2" = "Load Game", "3" = "Exit"
	// choice := "2"

		choice := "1"

		switch choice{
			case "1": 
				fmt.Println("Start Game")
			case "2": 
				fmt.Println("Load Game")
			case "3": 
				fmt.Println("Exit")
		}

}