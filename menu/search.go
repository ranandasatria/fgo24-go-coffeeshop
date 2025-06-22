package menu

import (
	"fgo24-go-weeklytask/utils"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func SearchFoods() {
	var message string
	cartManager := &CartManager{}

	for {
		utils.ClearScreen()
		if message != "" {
			fmt.Println(message)
			message = ""
		}

		fmt.Println("--- Search ---")
		fmt.Print("Enter keyword to search (or 'b' to return): ")
		var keyword string
		fmt.Scanln(&keyword)

		if keyword == "b" {
			MainMenu()
			return
		}

		keyword = strings.ToLower(keyword)
		var results []*Food
		var mu sync.Mutex
		var wg sync.WaitGroup

		for i := range FoodList {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				if strings.Contains(strings.ToLower(FoodList[index].Name), keyword) {
					mu.Lock()
					results = append(results, &FoodList[index])
					mu.Unlock()
				}
			}(i)
		}

		wg.Wait()

		if len(results) == 0 {
			message = "❌ No matching items found."
			continue
		}

		fmt.Println("\n--- Search Results ---")
		for i, item := range results {
			fmt.Printf("%d. %s: Rp%d\n", i+1, item.Name, item.Price)
		}
		fmt.Println("b. Back")

		choice, err := utils.ReadStringInput("Select item to add to cart: ")
		if err != nil {
			message = "Invalid choice."
			continue
		}

		if choice == "b" {
			MainMenu()
			return
		}

		choiceNum, err := strconv.Atoi(choice)
		if err != nil {
			message = "Invalid choice."
			continue
		}

		if choiceNum > 0 && choiceNum <= len(results) {
			if cartManager.AddToCart(results[choiceNum-1]) {
				message = "Item added to cart."
			} else {
				message = "Cancelled."
			}
			continue
		}

		message = "Invalid choice."
	}
}