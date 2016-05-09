package main

import "fmt"
import "time"

func BubbleSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		flag := true
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag := false
			}
		}
		if flag == false {
			break
		}
	}

}

func main() {
	values := []int{1, 5, 2, 3, 9, 42, 12}
	BubbleSort(values)
	fmt.Println(values)

}
