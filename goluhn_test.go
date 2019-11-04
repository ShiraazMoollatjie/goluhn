package goluhn

import (
	"log"
	"strconv"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		number   string
		expectOK bool
	}{
		{"1234567812345670", true},
		{"1111222233334444", true},
		{"1111222233334441", false},
		{"49927398716", true},
		{"49927398717", false},
		{"1234567812345678", false},
		{"79927398710", false},
		{"79927398711", false},
		{"79927398712", false},
		{"79927398713", true},
		{"79927398714", false},
		{"79927398715", false},
		{"79927398716", false},
		{"79927398717", false},
		{"79927398718", false},
		{"79927398719", false},
		{"374652346956782346957823694857692364857368475368", true},
		{"374652346956782346957823694857692364857387456834", false},
		{"8", false},
		{"0", true},
	}
	for _, test := range tests {
		t.Run(test.number, func(t *testing.T) {
			res := Validate(test.number)
			if res != nil {
				if test.expectOK {
					log.Printf("Expected success but luhn check unsuccessful for %s.", test.number)
					t.Fail()
				}
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		number    string
		luhnDigit string
		expected  string
	}{
		{"123456781234567", "0", "1234567812345670"},
		{"111122223333444", "4", "1111222233334444"},
		{"7992739871", "3", "79927398713"},
		{"37465234695678234695782369485769236485736847536", "8", "374652346956782346957823694857692364857368475368"},
	}
	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			l, n, err := Calculate(test.number)
			if err != nil {
				log.Printf("Unexpected err %+v", err)
				t.Fail()
			}
			if test.luhnDigit != l {
				log.Printf("Expected luhn digit %s. Actual luhn digit %s", test.luhnDigit, l)
				t.Fail()
			}
			if n != test.expected {
				log.Printf("Expected %s to generate luhn number %s", test.number, test.expected)
				t.Fail()
			}
			err = Validate(n)
			if err != nil {
				log.Printf("Cannot validate derive luhn number %s. Error: %+v", n, err)
				t.Fail()
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		numberSize int
		sampleSize int
	}{
		{1, 100},
		{10, 1000},
		{100, 1000},
		{1000, 1000},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.numberSize), func(t *testing.T) {
			for i := 0; i < test.sampleSize; i++ {
				err := Validate(Generate(test.numberSize))
				if err != nil {
					log.Printf("Unexpected err %+v", err)
					t.Fail()
				}
			}
		})
	}
}
