package menu

import (
	"fgo24-go-weeklytask/utils"
	"fmt"
)

var Cart []*Food

func AddToCart(food *Food) bool {
	var input string
	fmt.Printf("Add '%s' to cart? (y/n): ", food.Name)
	fmt.Scanln(&input)

	if input == "y" || input == "Y" {
		Cart = append(Cart, food)
		return true
	}
	return false
}

func ShowCart() {
	utils.ClearScreen()

	if len(Cart) == 0 {
		for {
			fmt.Println("Cart is empty.")
			fmt.Println("1. Order\n9. Back")

			var choice int
			fmt.Print("Input choice: ")
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				ShowFoodCategory()
				return
			case 9:
				MainMenu()
				return
			default:
				fmt.Println("Invalid choice.")
			}
		}
	}

	fmt.Println("--- Cart ---")
	for i, item := range Cart {
		fmt.Printf("%d. %s: Rp%d\n", i+1, item.Name, item.Price)
	}
	fmt.Println("\n9. Back")

	for {
		var choice int
		fmt.Print("Input choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 9:
			MainMenu()
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
