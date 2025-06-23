package menu

import (
	"fgo24-go-weeklytask/controllers"
	"fgo24-go-weeklytask/utils"
	"fmt"
)

func ShowCart() {
	var message string

	for {
		utils.ClearScreen()
		if message != "" {
			fmt.Println(message)
			message = ""
		}

		if len(controllers.Cart) == 0 {
			fmt.Println("Cart is empty.")
			fmt.Println("o. Order")
			fmt.Println("b. Back")
		} else {
			fmt.Println("--- Cart ---")
			for i, item := range controllers.Cart {
				fmt.Printf("%d. %s: Rp%d\n", i+1, item.Name, item.Price)
			}
			fmt.Println("\nb. Back")
		}

		input, err := utils.ReadStringInput("Input choice: ")
		if err != nil {
			message = "Invalid input."
			continue
		}

		if len(controllers.Cart) == 0 {
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
