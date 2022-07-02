package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}
}

func fizzBuzz(number int) string {
	var result string

	isFizz := (number % 3) == 0
	isBuzz := (number % 5) == 0

	if isFizz {
		result = "Fizz"
	}

	if isBuzz {
		if isFizz {
			result = "Fizz Buzz"
		} else {
			result = "Buzz"
		}
	}

	if len(result) == 0 {
		result = strconv.Itoa(number)
	}

	return result
}
