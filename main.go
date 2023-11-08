package main

import "fmt"

func IsPalindrome(input int) bool {
	i := 1
	newNumber := 0
	for input%i != input {
		newNumber = (newNumber * 10) + ((input / i) % 10)
		i = i * 10
	}
	return (newNumber == input)

}

// reverse
// equate

// split digits using modulus
// write digits in reverse order
// equate

func main() {
	fmt.Println(IsPalindrome(123)) // false
	fmt.Println(IsPalindrome(121)) // true
	fmt.Println(IsPalindrome(1))   // true
	fmt.Println(IsPalindrome(11))  // true
	fmt.Println(IsPalindrome(12))  // false
}
