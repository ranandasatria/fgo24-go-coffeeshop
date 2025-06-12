package utils

import (
	"fmt"
	"strings"
)

func GetIntInput(prompt string) (int, error) {
	fmt.Print(prompt)
	var input int
	_, err := fmt.Scanln(&input)
	return input, err
}

func GetStringInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scanln(&input)
	return strings.ToUpper(strings.TrimSpace(input)), err
}