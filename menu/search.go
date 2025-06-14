package menu

import (
	"fmt"
	"strings"
)

func SearchFoods() {
	var keyword string
	fmt.Print("Enter keyword to search: ")
	fmt.Scanln(&keyword)

	keyword = strings.ToLower(keyword)

	var results []*Food
	for i := range FoodList {
		if strings.Contains(strings.ToLower(FoodList[i].Name), keyword) {
			results = append(results, &FoodList[i])
		}
	}

	if len(results) == 0 {
		fmt.Println("❌ No matching items found.")
	} else {
		fmt.Println("--- Search Results ---")
		for i, item := range results {
			fmt.Printf("%d. %s: Rp%d\n", i+1, item.Name, item.Price)
		}

		var choice int
		fmt.Print("Select item to add to cart (0 to cancel): ")
		fmt.Scanln(&choice)

		if choice > 0 && choice <= len(results) {
			if AddToCart(results[choice-1]) {
				fmt.Println("Item added to cart.")
				MainMenu()
				return
			} else {
				fmt.Println("Cancelled.")
				MainMenu()
				return
			}
		} else {
			fmt.Println("No item added.")
			MainMenu()
				return
		}
	}
}
