package main

import "fmt"

func main() {
	var n uint16
	var sum int32
	var i uint16

	fmt.Scan(&n)

	s := make([]int32, n)

	for i = 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	for i = 0; i < n; i++ {
		sum = sum + s[i]
	}

	fmt.Print(sum)
}
