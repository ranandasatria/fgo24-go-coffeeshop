package menu

import (
	"fgo24-go-weeklytask/utils"
	"fmt"
)

func ShowFoodsByCategory(category string) {
	var message string

	for {
		utils.ClearScreen()
		if message != "" {
			fmt.Println(message)
			message = ""
		}

		fmt.Printf("--- %s ---\n", category)

		var availableItems []*Food

		for i := range FoodList {
			if FoodList[i].Category == category {
				availableItems = append(availableItems, &FoodList[i])
				fmt.Printf("%d. %s: Rp%d\n", len(availableItems), FoodList[i].Name, FoodList[i].Price)
			}
		}
		fmt.Println("9. Back")

		var choice int
		fmt.Print("Input choice: ")
		fmt.Scanln(&choice)

		if choice > 0 && choice <= len(availableItems) {
			success := AddToCart(availableItems[choice-1])
			if success {
			message = "Item added to cart."
			} else {
				message = "Cancelled."
			}
		} else if choice == 9 {
			ShowFoodCategory()
		} else {
			message = "Invalid choice."
		}
	}
}
