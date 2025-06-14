package menu

import (
	"fgo24-go-weeklytask/utils"
	"fmt"
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
	utils.ClearScreen()

	fmt.Println("--- Food Category ---")
	for _, item := range FoodCategoryMenu {
		fmt.Println(item)
	}

	var input int
	fmt.Print("Input choice: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		ShowFoodsByCategory("Meals")
	case 2:
		ShowFoodsByCategory("Drinks")
	case 3:
		ShowFoodsByCategory("Snacks")
	case 4:
		ShowFoodsByCategory("Desserts")
	case 9:
		MainMenu()
	default:
		fmt.Println("Invalid choice.")
	}

}
