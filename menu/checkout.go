package menu

import (
	"fgo24-go-weeklytask/utils"
	"fmt"
)

func Checkout() {
	if len(Cart) == 0 {
		MainMenu("Cart is empty. Nothing to checkout.")
		return
	}

	for {
		utils.ClearScreen()
		fmt.Println("--- Checkout ---")
		printCartItems()

		fmt.Print("Proceed to checkout? (y/n): ")
		var input string
		fmt.Scanln(&input)

		if input == "y" || input == "Y" {
			Cart = []*Food{}
			MainMenu("✅ Checkout successful! Thank you for your purchase.")
			return
		} else if input == "n" || input == "N" {
			MainMenu("Checkout canceled.")
			return
		} else {
			fmt.Println("Invalid input. Please enter y or n.")
		}
	}
}

func printCartItems() {
	for i, item := range Cart {
		fmt.Printf("%d. %s: Rp%d\n", i+1, item.Name, item.Price)
	}
}


