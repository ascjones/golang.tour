package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func SqrtWithErr(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z, delta := float64(1), float64(1)
	for math.Abs(delta) > 0.001 {
		s := z - (z*z-x)/(2*z)
		delta = z - s
		z = s
	}
	return z, nil
}

func main() {
	fmt.Println(SqrtWithErr(2))
	fmt.Println(SqrtWithErr(-2))
}
