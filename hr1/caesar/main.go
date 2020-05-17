package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret1 := ""
	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetLower, ch) >= 0:
			ret1 = ret1 + string(rotate(ch, delta, []rune(alphabetLower)))
		case strings.IndexRune(alphabetUpper, ch) >= 0:
			ret1 = ret1 + string(rotate(ch, delta, []rune(alphabetUpper)))
		default:
			ret1 = ret1 + string(ch)
		}
	}
	fmt.Println("Solution 1 gives:", ret1)

	var ret2 []rune
	for _, ch := range input {
		ret2 = append(ret2, cipher(ch, delta))
	}
	fmt.Println("Solution 2 gives:", string(ret2))
}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r >= 'Z' {
		return rotateWithBase(r, 'A', delta)
	}
	if r >= 'a' && r <= 'z' {
		return rotateWithBase(r, 'a', delta)
	}
	return r
}

func rotateWithBase(r rune, base, delta int) rune {
	tmp := int(r) - 'A'
	tmp = (tmp + delta) % base
	return rune(tmp + 'A')
}

func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("idx < 0")
	}
	idx = (idx + delta) % len(key)
	return key[idx]
}
