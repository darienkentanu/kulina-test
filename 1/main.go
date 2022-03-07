package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(kulinaFood(15))
}

func kulinaFood(angka int) (output []string) {
	for i := 1; i <= angka; i++ {
		if i%3 == 0 && i%5 == 0 {
			output = append(output, "Kulina x Food")
		} else if i%3 == 0 {
			output = append(output, "Kulina")
		} else if i%5 == 0 {
			output = append(output, "Food")
		} else {
			output = append(output, strconv.Itoa(i))
		}
	}
	return output
}
