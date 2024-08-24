package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan string untuk divalidasi: ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	isValidSymbol := ValidateSymbol(input)
	fmt.Println("Output:", isValidSymbol)
}

func ValidateSymbol(s string) bool {
	stack := []rune{}
	matchingSymbol := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{', '<':
			stack = append(stack, char)
		case ')', ']', '}', '>':
			if len(stack) == 0 || stack[len(stack)-1] != matchingSymbol[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
