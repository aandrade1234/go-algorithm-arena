package main

import (
	"errors"
	"fmt"
	"math/big"
)

var ErrNegativeInput = errors.New("factorial is not defined for negative numbers")

func main() {

	fmt.Println(findFactorialInteractive(-1))
	fmt.Println(findFactorialInteractive(0))
	fmt.Println(findFactorialInteractive(10))
	fmt.Println()
	fmt.Println(findFactorialRecursive(-1))
	fmt.Println(findFactorialRecursive(0))
	fmt.Println(findFactorialRecursive(10))
}

func findFactorialInteractive(value int) (*big.Int, error) {
	if value < 0 {
		return nil, ErrNegativeInput
	}

	result := big.NewInt(1)
	for i := int64(2); i <= int64(value); i++ {
		result.Mul(result, big.NewInt(i))
	}

	return result, nil
}

func findFactorialRecursive(value int) (*big.Int, error) {
	switch {
	case value < 0:
		return nil, ErrNegativeInput
	case value < 1:
		return big.NewInt(1), nil
	default:
		val, err := findFactorialRecursive(value - 1)
		if err != nil {
			return nil, err
		}

		return val.Mul(val, big.NewInt(int64(value))), nil
	}
}
