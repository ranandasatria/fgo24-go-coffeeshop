package main

import (
	"fmt"
	"playground/utils"
)

func main() {
	// cart := utils.NewCart() 

	for {
		utils.DisplayMainMenu()
		choice, err := utils.GetIntInput("Masukkan pilihan: ")
		if err != nil {
			fmt.Println("Input tidak valid, coba lagi.")
			continue
		}

		if choice == 1 {
			for {
				utils.DisplayFoodMenu()
				foodChoice, err := utils.GetIntInput("Masukkan pilihan: ")
				if err != nil {
					fmt.Println("Input tidak valid, coba lagi.")
					continue
				}

				if foodChoice == 5 {
					break
				}

				if foodChoice < 1 || foodChoice > 4 {
					fmt.Println("Pilihan tidak valid!")
					continue
				}

				item, price, ok := utils.SelectFoodItem(foodChoice)
				if !ok {
					fmt.Println("Pilihan tidak valid!")
					continue
				}

				fmt.Printf("Apakah yakin memilih %s? (%s)\n", item, price)
				confirm, err := utils.GetStringInput("Jawab (YA/TIDAK): ")
				if err != nil {
					fmt.Println("Input tidak valid, coba lagi.")
					continue
				}

				if confirm == "YA" {
					// cart.AddItem(item, price)
					fmt.Println("Item ditambahkan ke keranjang!")
				} else {
					fmt.Println("Pesanan dibatalkan.")
				}
			}
		} else if choice == 2 {
			// cart.Display()
		} else if choice == 3 {
			fmt.Println("---")
			fmt.Println("Terima kasih telah memesan!")
			// cart.Display()
			return 
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}