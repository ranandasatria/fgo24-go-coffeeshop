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
		fmt.Print("Enter keyword to search (or 'back' to return): ")
		var keyword string
		fmt.Scanln(&keyword)

		if keyword == "back" {
			MainMenu()
			return
		}

		keyword = strings.ToLower(keyword)
		var results []*Food
		var mu sync.Mutex
		var wg sync.WaitGroup

		chunkSize := (len(FoodList) + 3) / 4 
		for i := 0; i < len(FoodList); i += chunkSize {
			end := i + chunkSize
			if end > len(FoodList) {
				end = len(FoodList)
			}

			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				localResults := []*Food{}
				for j := start; j < end; j++ {
					if strings.Contains(strings.ToLower(FoodList[j].Name), keyword) {
						localResults = append(localResults, &FoodList[j])
					}
				}
				mu.Lock()
				results = append(results, localResults...)
				mu.Unlock()
			}(i, end)
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