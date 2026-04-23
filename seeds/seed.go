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
		ImageURL:    "https://i.ibb.co.com/Myfj9xd7/nasi-goreng.jpg",
	},
	{
		Name:        "Nasi Magelangan",
		Price:       26000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Perpaduan nasi goreng dan mie dengan cita rasa gurih",
		ImageURL:    "https://i.ibb.co.com/nsRNqyLJ/nasi-goreng-mawut-magelangan-foto-resep-utama.jpg",
	},
	{
		Name:        "Nasi Omlet",
		Price:       23000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Nasi dengan telur omelet lembut dan saus spesial",
		ImageURL:    "https://i.ibb.co.com/RGz3V9wk/Cara-Membuat-Nasi-Omelet-Khas-Jepang-2.jpg",
	},
	{
		Name:        "Mie Dok-Dok",
		Price:       22000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Mie kuah khas dengan bumbu rempah hangat",
		ImageURL:    "https://i.ibb.co.com/xKJTpb0N/62146ddadba5a.jpg",
	},
	{
		Name:        "Mie Goreng",
		Price:       22000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Mie goreng dengan sayur dan topping pilihan",
		ImageURL:    "https://i.ibb.co.com/hJDWRMX7/mie-goreng-saus-tiram.jpg",
	},
	{
		Name:        "Mie Rebus",
		Price:       22000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Mie rebus hangat dengan kuah gurih",
		ImageURL:    "https://i.ibb.co.com/5XxmyDMF/cccb39c7ae48e92058f2f99bc36aacd8.webp",
	},
	{
		Name:        "Soto Sokaraja",
		Price:       28000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Soto khas Sokaraja dengan kuah gurih dan sambal kacang",
		ImageURL:    "https://i.ibb.co.com/8L7YdTdB/b83edf8f-acd7-4c19-be25-08bc31002ae7.jpg",
	},
	{
		Name:        "French Fries",
		Price:       18000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Kentang goreng renyah dengan saus pilihan",
		ImageURL:    "https://i.ibb.co.com/KjkF6BVj/50223-homemade-crispy-seasoned-french-fries-VAT-Beauty-4x3-789ecb2eaed34d6e879b9a93dd56a50a.jpg",
	},
	{
		Name:        "Roti Bakar",
		Price:       20000,
		Category:    "Makanan",
		Stock:       50,
		Description: "Roti bakar manis dengan berbagai topping",
		ImageURL:    "Roti Bakar",
	},
	{
		Name:        "Le Mineral",
		Price:       8000,
		Category:    "Minuman",
		Stock:       50,
		Description: "Air mineral botol",
		ImageURL:    "https://i.ibb.co.com/YB3gM6Wr/338234-18-9-2019-15-50-20-1665805161.webp",
	},
	{
		Name:        "Es Jeruk",
		Price:       12000,
		Category:    "Minuman",
		Stock:       50,
		Description: "Minuman jeruk segar dingin",
		ImageURL:    "https://i.ibb.co.com/fzF50QxJ/ES-JERUK-murni.jpg",
	},
	{
		Name:        "Es Teh Manis",
		Price:       10000,
		Category:    "Minuman",
		Stock:       50,
		Description: "Teh manis dingin segar",
		ImageURL:    "https://i.ibb.co.com/TxgCnwBY/cover-w480-h480-sku-teh-es.png",
	},
	{
		Name:        "Kopi Brazil",
		Price:       18000,
		Category:    "Minuman",
		Stock:       50,
		Description: "Kopi brazil dengan aroma khas",
		ImageURL:    "https://i.ibb.co.com/yn9G0jxB/Sprudge-Summer-Drinks-Brazil-Juliana-Ganan-Tropicalia-Cafe-Secreto-Cicero-Rodrigues-13-740x494.jpg",
	},

	}
	for _, p := range products {
	config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}