package main

import (
	"fmt"
	"strconv"
)

func fibonacci(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func main() {
	for i := 0; i <= 9; i++ {
		fmt.Print(strconv.Itoa(fibonacci(i)) + " ")
	}
}