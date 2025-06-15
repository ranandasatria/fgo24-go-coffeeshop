package menu

import (
	"fgo24-go-weeklytask/utils"
	"fmt"
	"os"
)

var FoodCategoryMenu = []string{
	"1. Meals",
	"2. Drinks",
	"3. Snacks",
	"4. Desserts",
	"9. Back",
	"0. Exit",
}

func ShowFoodCategory() {
	var message string

	for {
		utils.ClearScreen()
		if message != "" {
			fmt.Println(message)
			message = ""
		}

		fmt.Println("--- Food Category ---")
		for _, item := range FoodCategoryMenu {
			fmt.Println(item)
		}

		input, err := utils.ReadIntInput("Input choice: ")
		if err != nil {
			message = "Invalid choice."
			continue
		}

		switch input {
		case 1:
			ShowFoodsByCategory("Meals")
			return
		case 2:
			ShowFoodsByCategory("Drinks")
			return
		case 3:
			ShowFoodsByCategory("Snacks")
			return
		case 4:
			ShowFoodsByCategory("Desserts")
			return
		case 9:
			MainMenu()
			return
		case 0:
			fmt.Println("Thank you!")
			os.Exit(0)
		default:
			message = "Invalid choice."
		}
	}
}