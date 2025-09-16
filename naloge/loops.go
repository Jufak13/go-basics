package naloge

import (
	"fmt"
	"strings"
)

func NalogeIzLoops() {
	// 0. Izpiši elemente seznama
	l := []int{1, 2, 3, 4, 5}
	m := map[string]string{"Hello": "World", "name": "nejc"}

	for _, item := range l {
		fmt.Println(item)
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

	// 1. Use a for loop to print numbers from 1 to 10
	// No variables needed

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. Check if a number is even or odd using if-else
	num := 7

	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 3. Print all even numbers from 1 to 20 using a loop and condition
	// No variables needed

	for i := 0; i <= 20; i += 2 {
		if i != 0 {
			fmt.Println(i)
		}
	}

	// 4. Use a loop to calculate the sum of numbers from 1 to 100
	sum := 0

	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 5. Use a loop to count how many numbers from 1 to 50 are divisible by 3
	count := 0

	for i := 1; i <= 50; i++ {
		if i%3 == 0 {
			count += 1
		}
	}
	fmt.Println(count)

	// 6. Write a loop that prints "Hello" 5 times
	// No variables needed

	for i := 1; i <= 5; i++ {
		fmt.Println("Hello")
	}

	// 7. Print the multiplication table of 5 (from 5x1 to 5x10)
	// No variables needed

	for i := 1; i <= 10; i++ {
		fmt.Println(i * 5)
	}

	// 8. Use an if statement to check if a person is eligible to vote (age >= 18)
	age := 19

	if age >= 18 {
		fmt.Println("Eligible")
	} else {
		fmt.Println("NOT!! ELIGIBLEE")
	}

	// 9. Print numbers from 1 to 20, but skip even numbers using continue
	// No variables needed

	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}

	// 10. Sum only odd numbers between 1 and 100
	sum = 0

	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 11. Create a countdown from 10 to 1 using a loop
	// No variables needed

	for i := 10; i >= 0; i-- {
		fmt.Println(i)
		// time.Sleep(1 * time.Second)
	}
	fmt.Println("Boom!")

	// 12. Use a nested loop to print a 5x5 star grid
	// No variables needed

	visina := 3
	sirina := 3

	for vrstice := 0; vrstice < visina; vrstice++ {
		vrsta := ""
		for zvezdice := 0; zvezdice < sirina; zvezdice++ {
			vrsta = fmt.Sprintf("%s*", vrsta)
		}
		fmt.Println(vrsta)
	}

	// 13. Check if a number is positive, negative or zero
	num = -3

	if num < 0 {
		fmt.Println("Negative")
	} else if num > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("0")
	}

	// 14. Use a loop to find the factorial of a number
	n := 5
	factorial := 0

	for i := 1; i <= n; i++ {
		factorial += i
	}

	fmt.Println(factorial)

	// 15. Loop through numbers 1 to 30 and print "Fizz" if divisible by 3, "Buzz" if divisible by 5, and "FizzBuzz" if both
	// No variables needed

	for i := 1; i <= 30; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else {
			fmt.Println(i)
		}
	}

	// 16. Use a loop to find the largest number in a slice
	numbers := []int{4, 10, 2, 99, 23}
	največja := 0

	for _, item := range numbers {
		if item > največja {
			največja = item
		}
	}
	fmt.Println(največja)

	// 17. Ask the user for 5 test scores and calculate the average using a loop
	scores1 := []int{80, 90, 85, 75, 95}
	sum = 0

	for _, item := range scores1 {
		sum += item
	}
	average := sum / len(scores1)
	fmt.Println(average)

	// 18. Use an if-else to determine the grade based on score
	score := 83

	if score > 80 {
		fmt.Println("A")
	} else if score < 80 && score > 60 {
		fmt.Println("B")
	} else if score < 60 && score > 40 {
		fmt.Println("C")
	} else if score < 40 && score > 20 {
		fmt.Println("D")
	} else if score < 20 {
		fmt.Println("F")
	}

	// 19. Create a loop that doubles a number until it’s greater than 1000
	n = 2

	for {
		n *= 2
		fmt.Println(n)
		if n > 1000 {
			break
		}
	}

	// 20. Loop through a slice and print only the numbers greater than 50
	numbers = []int{23, 67, 45, 89, 10, 102}

	for _, item := range numbers {
		if item > 50 {
			fmt.Println(item)
		}
	}

	// 21. Use a loop to reverse a string
	input := "hello"
	reversed := ""

	crke := strings.Split(input, "")

	for i := len(crke) - 1; i >= 0; i-- {
		reversed = fmt.Sprintf("%s%s", reversed, crke[i])
	}

	fmt.Println(reversed)

	// 22. Create a simple login check using if-else
	username := "admin"
	password := "pass123"

	if username == "admin" && password == "pass123" {
		fmt.Println("Welcome")
	} else {
		fmt.Println("Spizdi intruder!")
	}

	// 23. Use a for loop with range to print all characters in a string
	input = "golang"
	razdeljen := strings.Split(input, "")

	for _, item := range razdeljen {
		fmt.Println(item)
	}

	// 24. Create a loop that finds the sum of digits of a number
	num = 1234
	sum = 0

	for {
		currentDigit := num % 10
		num /= 10
		sum += currentDigit
		if num == 0 {
			break
		}
	}

	fmt.Println(sum)

	// 25. Create a menu system using switch:
	// "a" = "Add Item", "b" = "Remove Item", "c" = "Checkout"
	choice := "b"

	switch choice {
	case "a":
		fmt.Println("Add Item")
	case "b":
		fmt.Println("Remove Item")
	case "c":
		fmt.Println("Čekavt")
	}

	// 26. Print all prime numbers between 1 and 50
	// No variables needed

	max := 50
	for st := 1; st <= max; st++ {
		prime := true

		for i := 2; i <= (st/2)+1; i++ {
			if st%i == 0 {
				prime = false
			}
		}
		if prime {
			fmt.Println(st)
		}
	}

	// 27. Find the smallest number in a slice

	numbers = []int{4, 10, 2, 99, 23}
	najnižja := numbers[0]

	for _, item := range numbers {
		if item < najnižja {
			najnižja = item
		}
	}
	fmt.Println(najnižja)

	// 28. Loop through numbers 1 to 100 and stop the loop if number is divisible by 77
	// No variables needed

	for i := 1; i <= 100; i++ {
		if i == 77 {
			break
		} else {
			fmt.Println(i)
		}
	}

	// 29. Use a loop to generate the first 10 numbers of the Fibonacci sequence
	ab := 0
	ba := 1

	fib := []int{ab, ba}

	for i := 1; i <= 10; i++ {
		fib = append(fib, fib[len(fib)-2]+fib[len(fib)-1])
	}
	fmt.Println(fib)

}
