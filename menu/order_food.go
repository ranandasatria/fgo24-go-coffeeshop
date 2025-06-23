package menu

import (
	"fgo24-go-weeklytask/utils"
	"fmt"
)

func ShowFoodCategory() {
	var message string

	for {
		utils.ClearScreen()
		if message != "" {
			fmt.Println(message)
			message = ""
		}

		fmt.Println("--- Food Category ---")
		fmt.Println("1. Meals")
		fmt.Println("2. Drinks")
		fmt.Println("3. Snacks")
		fmt.Println("4. Desserts")
		fmt.Println("b. Back")

		input, err := utils.ReadStringInput("Input choice: ")
		if err != nil {
			message = "Invalid input."
			continue
		}

		switch input {
		case "1":
			ShowFoodsByCategory("Meals")
			return
		case "2":
			ShowFoodsByCategory("Drinks")
			return
		case "3":
			ShowFoodsByCategory("Snacks")
			return
		case "4":
			ShowFoodsByCategory("Desserts")
			return
		case "b":
			MainMenu()
			return
		default:
			message = "Invalid input."
		}
	}
}