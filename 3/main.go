package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("XII"))
	fmt.Println(romanToInt("LIV"))
	fmt.Println(romanToInt("MMXXII"))
}

func romanToInt(value string) int {
	result := 0
	number := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romawi := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var i int
	for i < len(romawi) {
		for j := 0; j < len(value); j++ {
			if j < len(value)-1 {
				temp := fmt.Sprintf("%s%s", string(value[j]), string(value[j+1]))
				if temp == romawi[i] {
					result += number[i] - number[i-1] - number[i+1]
				}
			}
			if string(value[j]) == romawi[i] {
				result += number[i]
			}
		}
		i++
	}
	return result
}
