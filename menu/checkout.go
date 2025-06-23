package menu

import (
	"fgo24-go-weeklytask/models"
	"fgo24-go-weeklytask/utils"
	"fmt"
)

func (c *CartManager) Checkout() {
	var message string

	for {
		utils.ClearScreen()
		if message != "" {
			fmt.Println(message)
			message = ""
		}

		if len(Cart) == 0 {
			MainMenu("Cart is empty. Nothing to checkout.")
			return
		}

		fmt.Println("--- Checkout ---")
		c.printCartItems()

		fmt.Print("Proceed to checkout? (y/n): ")
		var input string
		fmt.Scanln(&input)

		if input == "y" || input == "Y" {
			Cart = []*models.Food{}
			MainMenu("✅ Checkout successful! Thank you for your purchase.")
			return
		} else if input == "n" || input == "N" {
			MainMenu("Checkout canceled.")
			return
		} else {
			message = "Invalid input."
		}
	}
}

func (c *CartManager) printCartItems() {
	for i, item := range Cart {
		fmt.Printf("%d. %s: Rp%d\n", i+1, item.Name, item.Price)
	}
}