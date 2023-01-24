package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	result := make([]string, 0)

	var input string

	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &input)
		result = append(result, process(input))
	}

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}

}

func process(input string) string {
	if len(input) <= 10 {
		return input
	}
	output := string(input[0])
	temp := strconv.Itoa(len(input) - 2)
	output += temp
	output += string(input[len(input)-1])
	return output
}
