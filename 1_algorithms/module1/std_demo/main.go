package main

import (
	"fmt"

	module01 "example.com/v1/module1"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		gcd := module01.GCD(a, b)
		fmt.Println(gcd)
	}
}
