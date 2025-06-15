package utils

import (
	"fmt"
	"strings"
)

func ReadStringInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scanln(&input)
	return strings.TrimSpace(input), err
}