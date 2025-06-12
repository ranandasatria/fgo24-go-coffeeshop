package utils

import "fmt"

func DisplayMainMenu() {
	fmt.Println("---")
	fmt.Println("1. Pilih menu")
	fmt.Println("2. Lihat keranjang")
	fmt.Println("3. Checkout")
}

func DisplayFoodMenu() {
	fmt.Println("---")
	fmt.Println("1. Makanan")
	fmt.Println("2. Minuman")
	fmt.Println("3. Snack")
	fmt.Println("4. Dessert")
	fmt.Println("5. Kembali ke menu utama")
}

func SelectFoodItem(choice int) (string, string, bool) {
	if choice != 1 {
		return "", "", false
	}

	fmt.Println("---")
	fmt.Println("1. Nasi Goreng - Rp15.000")
	fmt.Println("2. Mie Goreng - Rp12.000")
	fmt.Println("3. Ayam Bakar - Rp20.000")
	itemChoice, err := GetIntInput("Masukkan pilihan: ")
	if err != nil {
		return "", "", false
	}

	if itemChoice == 1 {
		return "Nasi Goreng", "Rp15.000", true
	} else if itemChoice == 2 {
		return "Mie Goreng", "Rp12.000", true
	} else if itemChoice == 3 {
		return "Ayam Bakar", "Rp20.000", true
	}
	return "", "", false
}