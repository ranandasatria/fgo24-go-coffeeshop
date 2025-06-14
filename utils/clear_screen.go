package utils

import (
	"fmt"
)

func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}