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
	"0. Exit",
}

func MainMenu(message ...string) {
	utils.ClearScreen()

	if len(message) > 0 && message[0] != "" {
		fmt.Println(message[0])
	}

	fmt.Println("--- Home ---")
	for _, item := range mainmenu {
		fmt.Println(item)
	}

	var input int
	fmt.Print("Input choice: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		ShowFoodCategory()
	case 2:
		ShowCart()
	case 3:
	Checkout()
	case 4:
	SearchFoods()
	case 0:
		fmt.Println("Thank you!")
		os.Exit(0)
	default:
		MainMenu("Invalid choice.")
	}
}
