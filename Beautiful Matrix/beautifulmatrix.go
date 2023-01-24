package main

import (
	"fmt"
	"math"
)

func main() {

	var matrix [5][5]int
	pos := struct {
		x int
		y int
	}{}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scanf("%d", &matrix[i][j])
			if matrix[i][j] == 1 {
				pos.x = i
				pos.y = j
			}
		}

		fmt.Scanf("%d\n")

	}

	fmt.Println(math.Abs(float64(pos.x)-2) + math.Abs(float64(pos.y)-2))

}

// Very tricky I dunno
