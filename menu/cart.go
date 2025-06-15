package menu

import (
	"fgo24-go-weeklytask/utils"
	"fmt"
)

type CartManager struct{}

var Cart []*Food

func (c *CartManager) AddToCart(food *Food) bool {
	var input string
	fmt.Printf("Add '%s' to cart? (y/n): ", food.Name)
	fmt.Scanln(&input)

	if input == "y" || input == "Y" {
		Cart = append(Cart, food)
		return true
	}
	return false
}

func (c *CartManager) ShowCart() {
	var message string

	for {
		utils.ClearScreen()
		if message != "" {
			fmt.Println(message)
			message = ""
		}

		if len(Cart) == 0 {
			fmt.Println("Cart is empty.")
			fmt.Println("1. Order\n9. Back")
		} else {
			fmt.Println("--- Cart ---")
			for i, item := range Cart {
				fmt.Printf("%d. %s: Rp%d\n", i+1, item.Name, item.Price)
			}
			fmt.Println("\n9. Back")
		}

		choice, err := utils.ReadIntInput("Input choice: ")
		if err != nil {
			message = "Invalid choice."
			continue
		}

		if len(Cart) == 0 {
			switch choice {
			case 1:
				ShowFoodCategory()
				return
			case 9:
				MainMenu()
				return
			default:
				message = "Invalid choice."
			}
		} else {
			if choice == 9 {
				MainMenu()
				return
			}
			message = "Invalid choice."
		}
	}
}