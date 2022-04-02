package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := ReverseString(input)
	doubleRev, doubleRevErr := ReverseString(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %q\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %q\n", doubleRev, doubleRevErr)
}

func ReverseString(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}
