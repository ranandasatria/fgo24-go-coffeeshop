package menu

import (
	"fgo24-go-weeklytask/utils"
	"fmt"
	"os"
)

var mainmenu = []string{
	"1. Order",
	"2. Cart",
	"3. Checkout",
	"4. Search",
	"e. Exit",
}

func MainMenu(message ...string) {
	var msg string
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}

	var cartManager TransactionManager = NewCartManager()

	for {
		utils.ClearScreen()
		if msg != "" {
			fmt.Println(msg)
			msg = ""
		}

		fmt.Println("--- Home ---")
		for _, item := range mainmenu {
			fmt.Println(item)
		}

		input, err := utils.ReadStringInput("Input choice: ")
		if err != nil {
			msg = "Invalid input."
			continue
		}

		switch input {
		case "1":
			ShowFoodCategory()
			return
		case "2":
			cartManager.ShowCart()
			return
		case "3":
			cartManager.Checkout()
			return
		case "4":
			SearchFoods()
			return
		case "e":
			fmt.Println("Thank you!")
			os.Exit(0)
		default:
			msg = "Invalid input."
		}
	}
}
