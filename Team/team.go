package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	var count int = 0

	var a, b, c int

	for i := 0; i < n; i++ {
		fmt.Scanf("%d%d%d\n", &a, &b, &c)
		if a+b+c >= 2 {
			count++
		}
	}

	fmt.Println(count)

}
