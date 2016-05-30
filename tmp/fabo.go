package main

import (
	"fmt"
)

func fabo(i int) int {
	if i == 1 || i == 2 {
		return 1
	}
	sum := 1
	n := 1
	tmp := 0
	for k := 3; k <= i; k++ {
		tmp = sum
		sum += n
		n = tmp
	}
	return sum
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(fabo(i))
	}
}
