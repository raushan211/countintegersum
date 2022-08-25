package main

import "fmt"

func main() {
	num := 4
	fmt.Println(countEven(num))
}
func countEven(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		sum := digit(i)
		if sum%2 == 0 {
			count++
		}
	}
	return count
}

func digit(num int) int {
	sum := 0

	for num > 0 {
		rem := num % 10
		sum = sum + rem
		num = num / 10
	}
	return sum
}
