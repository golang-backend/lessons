package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	// balance := 100_000
	// fmt.Printf("Old Balance: %d\n", balance)
	// chargeBalance(&balance)

	// fmt.Println("Balance: ", balance)

	// Nested(chargeBalance, 10000)

	fn := func(a, b int) func(a, b int) int {
		return func(a, b int) int { return a + b }
	}

	func(a, b int) {
		fmt.Println(a + b)
	}(10, 34)

	fmt.Println(fn(23, 45))
}

func add(a, b int) int {
	return a + b
}

func chargeBalance(balance *int) error {
	// checking balance for negative number
	isNegativ := func(b int) error {
		if b < 0 {
			return fmt.Errorf("Negativ balance")
		}
		return nil
	}(*balance)

	if isNegativ != nil {
		return isNegativ
	}

	*balance = *balance - 10_000
	return nil
}

func Nested(ch func(*int) error, amount int) func() {
	if ch != nil {
		return func() {}
	}

	return func() {}
}
