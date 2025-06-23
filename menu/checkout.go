package menu

import (
	"fgo24-go-weeklytask/controllers"
	"fgo24-go-weeklytask/utils"
	"fmt"
	"time"
)

func Checkout() {
	var message string
	cartManager := controllers.NewCartManager()

	for {
		utils.ClearScreen()
		if message != "" {
			fmt.Println(message)
			message = ""
		}

		if len(controllers.Cart) == 0 {
			fmt.Println("Cart is empty. Add some items first")
			fmt.Println("b. Back")
		} else {
			fmt.Println("--- Checkout ---")
			total := 0
			for i, item := range controllers.Cart {
				fmt.Printf("%d. %s: Rp%d\n", i+1, item.Name, item.Price)
				total += item.Price
			}
			fmt.Printf("\nTotal: Rp%d\n", total)
			fmt.Println("y. Proceed")
			fmt.Println("b. Back")
		}

		input, err := utils.ReadStringInput("Input choice: ")
		if err != nil {
			message = "Invalid choice."
			continue
		}

		if input == "b" {
			MainMenu()
			return
		} else if input == "y" && len(controllers.Cart) > 0 {
			cartManager.ProceedCheckout()
			utils.ClearScreen()
			fmt.Println("✅ Checkout successful! Thank you for your purchase.")
			time.Sleep(2 * time.Second)
			MainMenu()
			return
		} else {
			message = "Invalid choice."
		}
	}
}