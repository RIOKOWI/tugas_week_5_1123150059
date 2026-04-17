package main

import (
	"log"

	"github.com/RIOKOWI/tugas_week_5_1123150059/config"
	"github.com/RIOKOWI/tugas_week_5_1123150059/models"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	config.InitDatabase()
	products := []models.Product{
		{
		Name:        "Nasi Goreng",
		Price:       25000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Nasi goreng dengan bumbu khas dan pelengkap",
		ImageURL:    "https://example.com/nasi-goreng.jpg",
	},
	{
		Name:        "Nasi Magelangan",
		Price:       26000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Perpaduan nasi goreng dan mie dengan cita rasa gurih",
		ImageURL:    "https://example.com/nasi-magelangan.jpg",
	},
	{
		Name:        "Nasi Omlet",
		Price:       23000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Nasi dengan telur omelet lembut dan saus spesial",
		ImageURL:    "https://example.com/nasi-omlet.jpg",
	},
	{
		Name:        "Mie Dok-Dok",
		Price:       22000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Mie kuah khas dengan bumbu rempah hangat",
		ImageURL:    "https://example.com/mie-dokdok.jpg",
	},
	{
		Name:        "Mie Goreng",
		Price:       22000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Mie goreng dengan sayur dan topping pilihan",
		ImageURL:    "https://example.com/mie-goreng.jpg",
	},
	{
		Name:        "Mie Rebus",
		Price:       22000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Mie rebus hangat dengan kuah gurih",
		ImageURL:    "https://example.com/mie-rebus.jpg",
	},
	{
		Name:        "Soto Sokaraja",
		Price:       28000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Soto khas Sokaraja dengan kuah gurih dan sambal kacang",
		ImageURL:    "https://example.com/soto-sokaraja.jpg",
	},
	{
		Name:        "French Fries",
		Price:       18000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Kentang goreng renyah dengan saus pilihan",
		ImageURL:    "https://example.com/french-fries.jpg",
	},
	{
		Name:        "Roti Bakar",
		Price:       20000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Roti bakar manis dengan berbagai topping",
		ImageURL:    "https://example.com/roti-bakar.jpg",
	},
	{
		Name:        "Le Mineral",
		Price:       8000,
		Category:    "Minuman",
		Stock:       50,
		Description: "Air mineral botol",
		ImageURL:    "https://example.com/le-mineral.jpg",
	},
	{
		Name:        "Es Jeruk",
		Price:       12000,
		Category:    "Minuman",
		Stock:       50,
		Description: "Minuman jeruk segar dingin",
		ImageURL:    "https://example.com/es-jeruk.jpg",
	},
	{
		Name:        "Es Teh Manis",
		Price:       10000,
		Category:    "Minuman",
		Stock:       50,
		Description: "Teh manis dingin segar",
		ImageURL:    "https://example.com/es-teh-manis.jpg",
	},
	{
		Name:        "Kopi Brazil",
		Price:       18000,
		Category:    "Minuman",
		Stock:       50,
		Description: "Kopi brazil dengan aroma khas",
		ImageURL:    "https://example.com/kopi-brazil.jpg",
	},

	}
	for _, p := range products {
	config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}