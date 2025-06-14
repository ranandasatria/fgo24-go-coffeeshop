package menu

type Food struct {
	Name     string
	Price    int
	Category string
}

var FoodList = []Food{
	{Name: "Nasi Goreng Kampung", Price: 15000, Category: "Meals"},
	{Name: "Mie Ayam Wonogiri", Price: 14000, Category: "Meals"},
	{Name: "Soto Betawi", Price: 20000, Category: "Meals"},
	{Name: "Ketoprak Cirebon", Price: 12000, Category: "Meals"},
	{Name: "Bakso Malang", Price: 17000, Category: "Meals"},

	{Name: "Es Teh Kampul", Price: 4000, Category: "Drinks"},
	{Name: "Es Jeruk Liar", Price: 5000, Category: "Drinks"},
	{Name: "Jus Terong Belanda", Price: 8000, Category: "Drinks"},
	{Name: "Air Kelapa", Price: 6000, Category: "Drinks"},
	{Name: "Air Putih Anget", Price: 1000, Category: "Drinks"},

	{Name: "Kacang Atom", Price: 3000, Category: "Snacks"},
	{Name: "Kerupuk Jangek", Price: 5000, Category: "Snacks"},
	{Name: "Otak-otak", Price: 2000, Category: "Snacks"},
	{Name: "Peyek", Price: 4000, Category: "Snacks"},
	{Name: "Gorengan", Price: 1500, Category: "Snacks"},

	{Name: "Es Krim Xue", Price: 19000, Category: "Desserts"},
	{Name: "Cendol Dawet", Price: 15000, Category: "Desserts"},
	{Name: "Martabak Mesir", Price: 20000, Category: "Desserts"},
	{Name: "Coklat Dubai", Price: 60000, Category: "Desserts"},
	{Name: "Mango Sago", Price: 30000, Category: "Desserts"},
}
