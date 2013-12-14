package main

import (
	"fmt"
	"math"
)

func newton(x, z float64) float64 {
	return z - (z*z-x)/(2*z)
}

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z = newton(x, z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	z, delta := float64(1), float64(1)
	i := 0
	for math.Abs(delta) > 0.001 {
		z2 := newton(x, z)
		delta = z - z2
		z = z2
		i++
	}
	fmt.Println(i, "iterations")
	return z
}

func main() {
	fmt.Println(Sqrt(2), Sqrt2(2), math.Sqrt(2))
	fmt.Println(Sqrt(3), Sqrt2(3), math.Sqrt(3))
}
