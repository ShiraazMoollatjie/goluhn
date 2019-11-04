package goluhn

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	asciiZero = 48
	asciiTen  = 57
)

// Validate returns an error if the provided string does not pass the luhn check.
func Validate(number string) error {
	p := len(number) % 2
	sum, err := calculateLuhnSum(number, p)
	if err != nil {
		return err
	}

	// If the total modulo 10 is not equal to 0, then the number is invalid.
	if sum%10 != 0 {
		return errors.New("invalid number")
	}

	return nil
}

// Calculate returns luhn check digit and the provided string number with its luhn check digit appended.
func Calculate(number string) (string, string, error) {
	p := (len(number) + 1) % 2
	sum, err := calculateLuhnSum(number, p)
	if err != nil {
		return "", "", nil
	}

	luhn := sum % 10
	if luhn != 0 {
		luhn = 10 - luhn
	}

	// If the total modulo 10 is not equal to 0, then the number is invalid.
	return strconv.FormatInt(luhn, 10), fmt.Sprintf("%s%d", number, luhn), nil
}

func calculateLuhnSum(number string, parity int) (int64, error) {
	var sum int64
	for i, d := range number {
		if d < asciiZero || d > asciiTen {
			return 0, errors.New("invalid digit")
		}

		d = d - asciiZero
		// Double the value of every second digit.
		if i%2 == parity {
			d *= 2
			// If the result of this doubling operation is greater than 9.
			if d > 9 {
				// The same final result can be found by subtracting 9 from that result.
				d -= 9
			}
		}

		// Take the sum of all the digits.
		sum += int64(d)
	}

	return sum, nil
}
