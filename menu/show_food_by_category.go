package menu

import (
	"fgo24-go-weeklytask/models"
	"fgo24-go-weeklytask/utils"
	"fmt"
	"strconv"
	"strings"
)

const itemsPerPage = 5

func ShowFoodsByCategory(category string) {
	var message string
	cartManager := &CartManager{}
	currentPage := 1

	for {
		utils.ClearScreen()
		if message != "" {
			fmt.Println(message)
			message = ""
		}

		fmt.Printf("--- %s (Page %d) ---\n", category, currentPage)

		var availableItems []*models.Food
		for i := range models.FoodList {
			if models.FoodList[i].Category == category {
				availableItems = append(availableItems, &models.FoodList[i])
			}
		}

		totalPages := (len(availableItems) + itemsPerPage - 1) / itemsPerPage
		if totalPages == 0 {
			totalPages = 1
		}

		startIdx := (currentPage - 1) * itemsPerPage
		endIdx := startIdx + itemsPerPage
		if endIdx > len(availableItems) {
			endIdx = len(availableItems)
		}

		for i := startIdx; i < endIdx; i++ {
			fmt.Printf("%d. %s: Rp%d\n", i+1, availableItems[i].Name, availableItems[i].Price)
		}

		navOptions := []string{}
		if currentPage > 1 {
			navOptions = append(navOptions, "[p. Previous]")
		}
		if currentPage < totalPages {
			navOptions = append(navOptions, "[n. Next]")
		}
		navOptions = append(navOptions, "[c. Cancel]")
		fmt.Printf("\nNavigation: %s\n", strings.Join(navOptions, " "))

		input, err := utils.ReadStringInput("Input choice: ")
		if err != nil {
			message = "Invalid input."
			continue
		}

		switch input {
		case "p":
			if currentPage > 1 {
				currentPage--
			} else {
				message = "No previous page."
			}
			continue
		case "n":
			if currentPage < totalPages {
				currentPage++
			} else {
				message = "No next page."
			}
			continue
		case "c":
			ShowFoodCategory()
			return
		}

		choice, err := strconv.Atoi(input)
		if err != nil {
			message = "Invalid input."
			continue
		}

		if choice >= startIdx+1 && choice <= endIdx {
			success := cartManager.AddToCart(availableItems[choice-1])
			if success {
				message = "Item added to cart."
			} else {
				message = "Cancelled."
			}
		} else {
			message = "Invalid input."
		}
	}
}
