package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	result := 0

	for i := 0; i < n; i++ {
		var input string
		fmt.Scanf("%s\n", &input)
		result += analyze(input)
	}

	fmt.Println(result)

}

func analyze(input string) int {

	if strings.Index(input, "+") != -1 {
		return 1
	} else {
		return -1
	}

}
