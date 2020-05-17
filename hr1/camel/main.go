package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	answer1 := solution1(input)
	fmt.Println("Solution 1 gives:", answer1)

	answer2 := solution2(input)
	fmt.Println("Solution 2 gives:", answer2)
}

func solution1(input string) int {
	answer := 1
	for _, ch := range input {
		min, max := 'A', 'Z'
		if ch >= min && max >= ch {
			answer++
		}
	}
	return answer
}
func solution2(input string) int {
	answer := 1
	for _, ch := range input {
		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}
	}
	return answer
}
