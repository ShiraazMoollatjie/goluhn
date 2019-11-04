# goluhn: A luhn number generator, validator and calculator in Go

> The Luhn algorithm or Luhn formula, also known as the "modulus 10" or "mod 10" algorithm, named after its creator, 
> IBM scientist Hans Peter Luhn, is a simple checksum formula used to validate a variety of identification numbers.

goluhn is a package that is able to:

* Validate a number to see if it is valid according to the luhn algorithm
* Calculate luhn check digits for a given number
* Generate random luhn numbers of any size

# How to use goluhn

## Validation

```go
err := goluhn.Validate("1111222233334444")
if res != nil {
  return err
}
```

## Calculation
The `Calculate` function returns the `luhnCheckDigit`, the `luhnNumber` (input number + check digit) or an `error`
```go
cd, n, err := goluhn.Calculate("111122223333444")
if res != nil {
  return err
}

fmt.Printf("Check digit: %s\n", cd)
fmt.Printf("Luhn number: %s\n", n)
```

# References

[Luhn algorithm on wikipedia](https://en.wikipedia.org/wiki/Luhn_algorithm)