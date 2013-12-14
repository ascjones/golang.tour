package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z, delta := complex128(1), complex128(1)
	for cmplx.Abs(delta) > 0.001 {
		s := z - (cmplx.Pow(z, 3)-x)/(3*cmplx.Pow(z, 2))
		delta = s - z
		z = s
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2), math.Cbrt(2))
}
