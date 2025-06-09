package main

import (
	"fmt"
	"playground/utils"
)

func main() {
	for {
		utils.DisplayMainMenu()
		choice, err := utils.GetIntInput("Masukkan pilihan: ")
		if err != nil {
			fmt.Println("Input tidak valid, coba lagi.")
		}

		if choice == 1 {
			fmt.Println("Ke menu makanan")
		} else if choice == 2 {
			fmt.Println("Ke menu keranjang")
		} else if choice == 3 {
			fmt.Println("Ke menu checkout")
		} else {
			fmt.Println("Input tidak valid, coba lagi.")
			return
		}
	}
}