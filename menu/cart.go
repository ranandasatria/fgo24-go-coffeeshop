package menu

import (
	"fgo24-go-weeklytask/models"
	"fgo24-go-weeklytask/utils"
	"fmt"
)

type TransactionManager interface {
	AddToCart(*models.Food) bool
	ShowCart()
	Checkout()
	printCartItems()
}

type CartManager struct {
	food []models.Food
}

var Cart []*models.Food

func NewCartManager() *CartManager {
	return &CartManager{
		food: []models.Food{},
	}
}

func (c *CartManager) AddToCart(food *models.Food) bool {
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
			fmt.Println("o. Order")
			fmt.Println("b. Back")
		} else {
			fmt.Println("--- Cart ---")
			for i, item := range Cart {
				fmt.Printf("%d. %s: Rp%d\n", i+1, item.Name, item.Price)
			}
			fmt.Println("\nb. Back")
		}

		input, err := utils.ReadStringInput("Input choice: ")
		if err != nil {
			message = "Invalid input."
			continue
		}

		if len(Cart) == 0 {
			switch input {
			case "o":
				ShowFoodCategory()
				return
			case "b":
				MainMenu()
				return
			default:
				message = "Invalid input."
			}
		} else {
			if input == "b" {
				MainMenu()
				return
			}
			message = "Invalid input."
		}
	}
}
