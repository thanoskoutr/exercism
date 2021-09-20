package luhn

import (
	"strconv"
	"strings"
)

func Valid(cardNumber string) bool {
	// Strip White Space
	cardNumberSplit := strings.Split(cardNumber, " ")
	cardNumberString := strings.Join(cardNumberSplit, "")

	// Strings of length 1 or less are not valid
	if len(cardNumberString) <= 1 {
		return false
	}

	// Slice to keep Credit Card numbers
	numbers := make([]int, len(cardNumberString))

	// All other non-digit characters are disallowed
	for i, char := range cardNumberString {
		val, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		numbers[i] = val
	}

	// Double every second digit, starting from the right
	for i := len(numbers) - 2; i >= 0; i -= 2 {
		numbers[i] *= 2
		if numbers[i] > 9 {
			numbers[i] -= 9
		}
	}

	// Then sum all of the digits
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	// If the sum is evenly divisible by 10, then the number is valid
	return sum%10 == 0
}
