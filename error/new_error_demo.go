package main

import (
	"fmt"
)

type ErrNegativeSqrt float64
type ErrSqrt struct {
	value       interface{}
	description string
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

func (e *ErrSqrt) Error() string {
	return fmt.Sprintf("%v:%v", e.description, e.value)
}

func Squrt(x interface{}) (interface{}, error) {
	var floatNum float64
	switch x.(type) {
	case int:
		floatNum = float64(x.(int))
	case float64:
		floatNum = x.(float64)
	default:
		return x, &ErrSqrt{value: x, description: "Wrong type"}
	}
	if floatNum < 0 {
		return floatNum, &ErrSqrt{value: floatNum, description: "Negative number"}
	}
	z, count := floatNum/2, 0
	for ; count < 10; count++ {
		z -= (z*z - floatNum) / (2 * z)
	}
	return z, nil
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z, count := x/2, 0
	for ; count < 10; count++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Squrt(2))
	fmt.Println(Squrt(-2))
	fmt.Println(Squrt("abc"))
}
