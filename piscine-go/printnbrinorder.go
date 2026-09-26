package piscine // declares the package name (must be "piscine" for 01-edu exercises)

import "github.com/01-edu/z01" // imports the z01 package used to print single runes

func PrintNbrInOrder(n int) { // function that prints digits of n in ascending order

	if n < 0 { // checks if the number is negative
		return // if negative, do nothing (negative numbers are excluded)
	}

	var digits [10]int // array to count occurrences of each digit (0 to 9)

	if n == 0 { // special case: if number is 0
		digits[0] = 1 // mark that digit 0 appears once
	}

	for n > 0 { // loop until all digits are processed
		d := n % 10 // extract the last digit of n
		digits[d]++ // increase count for that digit
		n = n / 10  // remove the last digit from n
	}

	for i := 0; i < 10; i++ { // loop from smallest digit (0) to largest (9)
		for digits[i] > 0 { // while this digit still has occurrences
			z01.PrintRune(rune(i + '0')) // convert digit to rune and print it
			digits[i]--                  // decrease count after printing
		}
	}
}
