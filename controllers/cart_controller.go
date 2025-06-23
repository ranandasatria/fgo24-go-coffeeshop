package controllers

import (
	"fgo24-go-weeklytask/models"
	"fmt"
)

type TransactionManager interface {
	AddToCart(*models.Food) bool
	ProceedCheckout()
}

type CartManager struct{}

var Cart []*models.Food

func NewCartManager() *CartManager {
	return &CartManager{}
}

func (c *CartManager) AddToCart(food *models.Food) bool {
	var input string
	fmt.Printf("Add '%s' to cart? (y/n): ", food.Name)
	fmt.Scanln(&input)

	if input == "y" || input == "Y" {
		Cart = append(Cart, food)
		return true
	}
	return false
}

func (c *CartManager) ProceedCheckout() {
	Cart = []*models.Food{}
}