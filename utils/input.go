package utils

import (
	"fmt"
)

func GetIntInput(prompt string) (int, error) {
	fmt.Print(prompt)
	var input int
	_, err := fmt.Scanln(&input)
	return input, err
}