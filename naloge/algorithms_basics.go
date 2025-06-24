package naloge

import (
	"fmt"
	"math"
	"strconv"
)

// 01 Return the smallest even number that appears more than once
// Test cases:
//   []int{2, 4, 2, 5, 6, 6} => 2
//   []int{3, 5, 7, 9} => "no repeating even values"
//   []int{} => "no repeating even values"
//   []int{10, 10, 8, 8, 6} => 8
//   []int{1, 2, 3, 2, 4} => 2

func SmallestRepeatingEven(numbers []int) (int, error) {
	frequencies := map[int]int{}
	for _, number := range numbers {
		if number%2 != 0 {
			continue
		}
		_, ok := frequencies[number] // ok = true, če je number obstaja v mapu
		if !ok {                     // če ni v mapu dodaj key in vrednost 1
			frequencies[number] = 1
		} else { // če je v mapu, prištej eno vrednosti key-a
			frequencies[number]++
		}
	}
	smallest := 0
	overriden := false
	for number, frequency := range frequencies {
		if !overriden && frequency > 1 {
			smallest = number
			overriden = true
			continue
		}

		if frequency > 1 && number < smallest {
			smallest = number
		}
	}

	if !overriden {
		return 0, fmt.Errorf("no repeating even values")
	}

	return smallest, nil
}

// 02 Return all numbers that are both prime and appear at least twice
// Test cases:
//   []int{2, 3, 3, 5, 5, 6} => [3 5]
//   []int{4, 6, 8, 9} => []
//   []int{2, 2, 2, 3, 3, 4, 4} => [2 3]
//   []int{1, 1, 1} => []
//   []int{} => []

func RepeatingPrimeNumbers(numbers []int) []int {
	primes := map[int]int{}
	out := []int{}

	for _, number := range numbers {
		if number == 1 {
			continue
		}

		isPrime := true
		for i := 2; i < number; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}

		if number == 2 {
			isPrime = true
		}

		if !isPrime {
			continue
		}

		_, ok := primes[number]
		if !ok {
			primes[number] = 1
		} else {
			primes[number] += 1
		}
	}

	for number, frequency := range primes {
		if frequency > 1 {
			out = append(out, number)
		}
	}

	return out
}

// 03 Return the average of all values that are at odd indices and divisible by 3
// Test cases:
//   []int{3, 6, 9, 12, 15, 18} => 12.0
//   []int{1, 2, 3, 4} => "no matching values at odd indices"
//   []int{} => "no matching values at odd indices"
//   []int{3, 3, 3, 3} => 3.0
//   []int{5, 6, 7, 9, 11, 12} => 9.0

// 1. preverim če je število iz seznama na lihem mestu v seznamu ter deljivo s 3
// 2. če je, potem ga dodam v seznam b
// 3. ko grem čez celotni prvi seznam a, zračunam povprečje števil v seznamu b ter vrnem rezultat
//
//
//

func AverageOddThrees(numbers []int) (float64, error) {
	average := []int{}
	result := 0
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no matching values at odd indices")
	}
	for i, number := range numbers {
		if i%2 != 0 && number%3 == 0 {
			average = append(average, number)
		}
	}
	if len(average) == 0 {
		return 0, fmt.Errorf("no matching values at odd indices")
	}

	for _, num := range average {
		result += num
	}
	result = result / len(average)

	return float64(result), nil
}

// 04 Return the number of unique values that are palindromes and odd
// Test cases:
//   []int{1, 3, 5, 7, 9, 11, 33, 44} => 7
//   []int{2, 4, 6, 8} => 0
//   []int{121, 131, 141, 151} => 4
//   []int{} => 0
//   []int{1, 1, 1, 1} => 1

// pali := strconv.Itoa(33)
// pali = "33"
// 1. Ustvarim nov seznam palindromes
// 2. preverim iz seznama če je število liho
// 3. če je, z funkcijo preverim oz. prepišem število v rikverc
// 4. primerjam številki in če sta enaki, dodam številko v seznam palindromes
// 5. na koncu seštejem koliko števil je v seznamu palindromes in vrnem rezultat

func OddPalindromes(numbers []int) int {
	palindromes := map[int]bool{}

	for _, number := range numbers {
		if number%2 != 0 && number > 9 {
			pali := strconv.Itoa(number)
			new := ""
			for i := len(pali) - 1; i >= 0; i-- {
				new += string(pali[i])
			}
			if pali == new {
				palindromes[number] = true
			}
		} else {
			if number%2 != 0 && number <= 9 {
				palindromes[number] = true
			}
		}
	}

	return len(palindromes)
}

// 05 Return the second largest value that appears more than once
// Test cases:
//   []int{5, 5, 6, 6, 7, 7} => 6
//   []int{1, 1, 2, 3, 4, 5} => "not enough repeating values"
//   []int{9, 9, 8, 8, 7, 7, 6, 6} => 8
//   []int{} => "not enough repeating values"
//   []int{2, 2, 2, 1, 1} => 1

// 1. ustvarim seznam dvojniki
// 2. preverim če je število dvojnik (grem čez seznam)
// 3. dodam dvojnike v seznam
// 4. razporedim svojnike naraščajoče (z algoritmom)
// 5. vrnem predzadnji element kot rezultat

func RepeatingSecondLargest(numbers []int) (int, error) {
	dvojniki := map[int]int{}
	seznam := []int{}

	for _, number := range numbers {
		dvojniki[number]++
	}
	for number, frequency := range dvojniki {
		if frequency > 1 {
			seznam = append(seznam, number)
		}
	}

	if len(seznam) < 2 {
		return 0, fmt.Errorf("not enough repeating values")
	}

	var highest int
	var secondHighest int

	if seznam[0] > seznam[1] {
		highest = seznam[0]
		secondHighest = seznam[1]
	} else {
		secondHighest = seznam[0]
		highest = seznam[1]
	}

	for _, number := range seznam[2:] {
		if number > highest {
			secondHighest = highest
			highest = number
		} else if number > secondHighest {
			secondHighest = number
		}
	}

	return secondHighest, nil
}

// 06 Return the longest increasing sub-slice from a list of integers
// Test cases:
//   []int{1, 2, 3, 2, 3, 4, 5} => [2 3 4 5]
//   []int{5, 4, 3, 2, 1} => [5]
//   []int{} => "empty input"
//   []int{1, 2, 1, 2, 3} => [1 2 3]
//   []int{10} => [10]

// 1. ustvaril bi 2 seznama.
// 2. šel bi čez dani seznam in preverjal, če številke naraščajo po vrsti.
// 3. če naraščajo, jih dodajam po vrsti v 1. seznam, dokler ne pridem do števila ki ne sledi zaporedju.
// 4. ko pridem do števila ki ni v zaporedju, grem v nov loop kjer ponovim preverjanje in dodajam številke v 2. seznam
// 5. ko pridem do števila ki ne sledi zaporedju preverim, če je len(2)>len(1), če je len(1) = len(2).
// 6. nadaljujem preverjanje in dodajam števila v 2. seznam.
// 7. ponavljam dokler ni konc seznama.
// 8. vrnem 1. seznam

func LongestIncreasingSubslice(numbers []int) ([]int, error) {
	if len(numbers) < 1 {
		return nil, fmt.Errorf("empty input")
	}

	seznam1 := []int{}
	seznam2 := []int{numbers[0]}
	num := numbers[0]

	for _, number := range numbers[1:] {
		if number > num {
			seznam2 = append(seznam2, number)
			num = number
		} else {
			if len(seznam2) > len(seznam1) {
				seznam1 = seznam2
				seznam2 = []int{number}
				num = number
			} else {
				seznam2 = []int{number}
				num = number
			}
		}

	}

	if len(seznam2) > len(seznam1) {
		return seznam2, nil
	}

	return seznam1, nil
}

// 07 Return the product of all odd numbers that are also perfect squares
// Test cases:
//   []int{1, 4, 9, 16, 25} => 225
//   []int{2, 3, 5, 7} => "no matching values"
//   []int{} => "no matching values"
//   []int{1, 1, 1} => 1
//   []int{49, 36, 25} => 1225

// 0. preverim če je prazen seznam
// 0.1 ustvarim spremenljivko "rezultat = 0"
// 1. grem čez seznam in preverjam število, če je liho
// 2. če je, preverim, če je deljiva sama s seboj (perf square)
// 3. prištejem k rezultatu
// 4. vrnem rezultat
// 5.
// 6.
// 7.
// 8.
// 9.
// 10.

func ProductOddPerfSquares(numbers []int) (int, error) {
	if len(numbers) < 1 {
		return 0, fmt.Errorf("no matching values")
	}

	rezultat := 1
	found := false

	for _, number := range numbers {
		if number%2 != 0 {
			rt := int(math.Sqrt(float64(number)))
			if rt*rt == number {
				rezultat *= number
				found = true
			}
		}

	}

	if !found {
		return 0, fmt.Errorf("no matching values")
	}

	return rezultat, nil
}

// 08 Return the first value that is both a multiple of 3 and appears exactly twice
// Test cases:
//   []int{3, 3, 6, 6, 9, 9} => 3
//   []int{2, 4, 6, 8, 10} => "no valid match found"
//   []int{} => "no valid match found"
//   []int{6, 6, 3, 3} => 6
//   []int{12, 12, 15} => 12

// 1. preverim če je seznam prazn in vrnem error če je
// 2. naredim mapo
// 3. preverim če je število deljivo z 3
// 4. če je ga dodam v mapo in dodam counter
// 5. ko zaključim preverjanje prvega seznama preverim v mapi katero število se prvo podvoji
// 6. vrnem to število
// 7.
// 8.
// 9.
// 10.

func FirstDoubleMultipleOfThree(numbers []int) (int, error) {
	if len(numbers) < 1 {
		return 0, fmt.Errorf("no valid match found")
	}

	frequencies := map[int]int{}

	for _, number := range numbers {
		frequencies[number]++
	}

	for _, number := range numbers {
		if number %3 != 0 {
			continue
		}

		value := frequencies[number]
		if value == 2 {
			return number, nil
		}
	}

	return 0, fmt.Errorf("no valid match found")
}

// 09 From two slices, return all values that are odd and only appear in one of the slices
// Test cases:
//   A: [1, 2, 3], B: [3, 4, 5] => [1 5]
//   A: [2, 4, 6], B: [1, 3, 5] => [1 3 5]
//   A: [1, 3], B: [1, 3] => []
//   A: [], B: [1, 2, 3] => [1 3]
//   A: [7, 9], B: [] => [7 9]

// 10 Return the index of the smallest number that appears more than once and is a multiple of 5
// Test cases:
//   []int{10, 20, 10, 30, 20, 40} => 0
//   []int{1, 2, 3, 4} => "no valid repeating multiple of 5"
//   []int{} => "no valid repeating multiple of 5"
//   []int{5, 5, 10, 10} => 0
//   []int{25, 30, 25, 30, 15, 15} => 4

// 11 Return the value that occurs most frequently, but only if it is a perfect square
// Test cases:
//   []int{4, 4, 9, 9, 9, 16} => 9
//   []int{2, 3, 5} => "no perfect square values"
//   []int{} => "no perfect square values"
//   []int{1, 1, 2, 2, 4, 4} => 1
//   []int{25, 36, 25, 36} => 25

// 12 Return the number of distinct primes that divide at least two different numbers in the list
// Test cases:
//   []int{6, 10, 15} => 2       // 2 and 5 divide multiple
//   []int{2, 3, 5, 7} => 0
//   []int{4, 6, 8, 10} => 2     // 2 and maybe 5
//   []int{} => 0
//   []int{9, 15, 21} => 1       // 3

// 13 Return the sum of the two most frequent even numbers that are also divisible by 4
// Test cases:
//   []int{4, 4, 8, 8, 12, 12} => "not enough matching values"
//   []int{2, 6, 10} => "not enough matching values"
//   []int{4, 4, 8} => "not enough matching values"
//   []int{16, 16, 20, 20, 24} => 36
//   []int{} => "not enough matching values"

// 14 Given a list of strings, return the shortest one that contains only digits and has no repeated characters
// Test cases:
//   []string{"123", "112", "789", "56"} => "56"
//   []string{"abc", "999", "111"} => "no valid string found"
//   []string{} => "no valid string found"
//   []string{"9876", "1234", "88"} => "88"
//   []string{"9", "98", "987"} => "9"

// 15 From a list of numbers, return the average of all values that:
// Test cases:
//   []int{3, 33, 6, 121, 303, 9} => 105.0     // 303 only one that passes
//   []int{3, 3, 6, 6, 9, 9} => "no matching values"
//   []int{6, 33, 303, 12321} => 169.5         // 33 and 303
//   []int{} => "no matching values"
//   []int{111, 222, 303, 303} => 111.0

// 16 Count how many groups of consecutive odd numbers sum to an even number
// Test cases:
//   []int{1, 3, 5, 2, 7, 9, 11} => 2      // [1 3 5] = 9 (odd), [7 9 11] = 27 (odd), skip; no even
//   []int{1, 3, 2, 5, 7, 2, 9, 11} => 1   // [5 7] = 12 (even)
//   []int{} => 0
//   []int{2, 4, 6} => 0
//   []int{1, 3, 5, 7} => 0

// 17 From a slice of strings, return the most common word after lowercasing and removing punctuation
// Test cases:
//   ["Hello!", "hello", "HELLO.", "Hi"] => "hello"
//   [] => "no input provided"
//   ["Yes?", "yes!", "YES", "no"] => "yes"
//   ["One", "Two", "Two", "Three."] => "two"
//   ["What's", "what's", "Whats"] => "whats"

// 18 From a slice of strings, return true if reversing all words and sorting gives the same as sorting and then reversing all
// Test cases:
//   ["abc", "def"] => true
//   ["abc", "cba"] => false
//   [] => true
//   ["a", "b", "c"] => true
//   ["xy", "yx"] => false

// 19 Given a list of integers, return the sum of values that:
// Test cases:
//   []int{2, 3, 5, 5, 7, 10} => 10
//   []int{4, 6, 8, 10} => "no qualifying values"
//   []int{} => "no qualifying values"
//   []int{13, 15, 17, 13} => 17
//   []int{1, 2, 3, 4} => "no qualifying values"

// 20 From a list of strings, return all that are palindromes after removing non-letter characters and lowercasing
// Test cases:
//   ["A man, a plan, a canal: Panama", "Racecar", "hello"] => ["A man, a plan, a canal: Panama", "Racecar"]
//   ["abc", "def"] => []
//   [] => []
//   ["No 'x' in Nixon"] => ["No 'x' in Nixon"]
//   ["Madam!", "madam", "MaDaM"] => ["Madam!", "madam", "MaDaM"]

// 21 Compute the average of all even non-negative numbers, rounded down.
// Test cases:
//   []int{2, 4, -1, -3, 6, 7} => 4
//   []int{1, 3, 5, -2} => "no even values"
//   []int{} => "no even values"
//   []int{-10, -20, -30} => "no even values"
//   []int{10, 20, 30} => 20

// 22 Return the longest string that starts and ends with the same letter (case-insensitive).
// Test cases:
//   []string{"Anna", "civic", "", "racecar", "apple"} => "racecar"
//   []string{"", "", ""} => "no valid word found"
//   []string{} => "no valid word found"
//   []string{"level", "stats", "bob"} => "stats"
//   []string{"wow", "deed", "deed"} => "deed"

// 23 Filter out all numbers less than or equal to 10.
// Find the smallest and largest of the remaining numbers.
// Count how many of the original numbers are strictly between those two.
// Test cases:
//   []int{5, 12, 15, 20, 25} => 2
//   []int{5, 8, 9} => "not enough values"
//   []int{11, 11, 11} => "not enough values"
//   []int{10, 20, 30} => 1
//   []int{50, 100, 70, 85} => 1

// 24 Return a list of strings that appear only once sorted by length.
// Test cases:
//   []string{"a", "bb", "ccc", "a", "bb", "dddd"} => ["ccc" "dddd"]
//   []string{"same", "same", "same"} => []
//   []string{} => []
//   []string{"one", "two", "three"} => ["one" "two" "three"]
//   []string{"a", "ab", "abc", "ab", "abc"} => ["a"]

// 25 Remove all non-positive numbers.
// Count how many digits each number has.
// Return the most common digit length.
// Test cases:
//   []int{1, 12, 123, 45, 678, 9, 10} => 2
//   []int{-1, -2, 0} => "no positive numbers"
//   []int{1000, 10000, 100000} => 5
//   []int{} => "no positive numbers"
//   []int{7, 88, 9, 101, 5} => 1

// 26 From a list of integers, collect all values divisible by 3.
// From those, remove all values that are also divisible by 5.
// Return them sorted in descending order.
// Test cases:
//   []int{3, 6, 10, 15, 18, 20} => [18 6 3]
//   []int{5, 10, 20, 25} => []
//   []int{} => []
//   []int{30, 45, 60} => []
//   []int{9, 12, 16} => [12 9]

// 27 From a list of words, remove all words shorter than 5 characters.
// Then, filter out any word that contains the letter 'z' or 'Z'.
// Return the remaining words sorted alphabetically.

// Test cases:
//   []string{"apple", "zebra", "banana", "Zoo", "plum"} => ["apple" "banana"]
//   []string{"zebra", "zoo", "zap"} => []
//   []string{} => []
//   []string{"grape", "melon", "berry"} => ["berry" "grape" "melon"]
//   []string{"aaazz", "bbbbb", "ccccz"} => ["bbbbb"]

// 28 From a list of integers, remove all duplicates.
// Then, split the numbers into two groups: even and odd.
// Return a slice containing first all even numbers, then all odd numbers, both sorted in ascending order.

// Test cases:
//   []int{1, 2, 3, 4, 5, 2, 3, 4} => [2 4 1 3 5]
//   []int{2, 4, 6, 8} => [2 4 6 8]
//   []int{1, 3, 5, 7} => [1 3 5 7]
//   []int{} => []
//   []int{5, 5, 5, 5} => [5]

// 29 From a list of positive integers, filter out values with an odd number of digits.
// From the rest, compute the product of the smallest and largest.
// Return the result as an integer.
// Test cases:
//   []int{10, 22, 333, 4444, 55555} => 220
//   []int{1, 12, 123, 1234, 12345} => 14808
//   []int{1, 3, 5, 7} => "no even-digit numbers"
//   []int{1000, 2000, 3000} => 3000000
//   []int{} => "no even-digit numbers"

// 30 From a list of words, remove all words that appear more than once.
// From the remaining words, collect the ones whose length is a prime number.
// Return them in reverse order of their appearance in the input.
// Test cases:
//   []string{"dog", "cat", "bird", "dog", "eagle"} => ["eagle" "cat"]
//   []string{"red", "red", "red"} => []
//   []string{} => []
//   []string{"hi", "there", "friend", "hi"} => ["friend" "there"]
//   []string{"prime", "seven", "nine", "ten"} => ["nine" "seven" "prime"]
