package menu

import (
	"fgo24-go-weeklytask/utils"
	"fmt"
	"strings"
)

func SearchFoods() {
	var message string
	cartManager := &CartManager{}

	for {
		utils.ClearScreen()
		if message != "" {
			fmt.Println(message)
			message = ""
		}

		fmt.Println("--- Search ---")
		fmt.Print("Enter keyword to search (or 'back' to return): ")
		var keyword string
		fmt.Scanln(&keyword)

		if keyword == "back" {
			MainMenu()
			return
		}

		keyword = strings.ToLower(keyword)
		var results []*Food
		for i := range FoodList {
			if strings.Contains(strings.ToLower(FoodList[i].Name), keyword) {
				results = append(results, &FoodList[i])
			}
		}

		if len(results) == 0 {
			message = "❌ No matching items found."
			continue
		}

		fmt.Println("\n--- Search Results ---")
		for i, item := range results {
			fmt.Printf("%d. %s: Rp%d\n", i+1, item.Name, item.Price)
		}
		fmt.Println("0. Cancel")

		choice, err := utils.ReadIntInput("Select item to add to cart (0 to cancel): ")
		if err != nil {
			message = "Invalid choice."
			continue
		}

		if choice == 0 {
			message = "No item added."
			continue
		}

		if choice > 0 && choice <= len(results) {
			if cartManager.AddToCart(results[choice-1]) {
				message = "Item added to cart."
			} else {
				message = "Cancelled."
			}
			continue
		}

		message = "Invalid choice."
	}
}