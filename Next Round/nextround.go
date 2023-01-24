package main

import "fmt"

func main() {

	var n, k int

	fmt.Scanf("%d%d\n", &n, &k)

	participants := make([]int, 0)

	for i := 0; i < n; i++ {
		var temp int
		fmt.Scanf("%d", &temp)
		participants = append(participants, temp)
	}

	if participants[0] == 0 {
		fmt.Println(0)
	} else {

		rank := 0
		result := 0

		for i := 0; i < n; i++ {
			if participants[i] == 0 {
				break
			}

			if i == 0 || participants[i] != participants[i-1] {
				rank = i + 1
			}

			if rank <= k {
				result++
			} else {
				break
			}

		}

		fmt.Println(result)
	}

}
