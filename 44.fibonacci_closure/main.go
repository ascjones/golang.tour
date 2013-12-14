package main

import "fmt"

func fibonacci2() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		if x == 1 && y == 1 {
			return 0
		}
		return x
	}
}

func main() {
	f := fibonacci2()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
