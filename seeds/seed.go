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
			Name:"Nasi Goreng Spesial", 
			Price:25000, 
			Category:"Makanan", 
			Stock:50,
			Description:"Nasi goreng dengan telur dan ayam",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"Teajus", 
			Price:5000, 
			Category:"Minuman", 
			Stock:50,
			Description:"Minuman Sachetan",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"Gudeg", 
			Price:20000, 
			Category:"Makanan", 
			Stock:50,
			Description:"Makanan Khas Yogyakarta",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"Kwetiaw", 
			Price:15000, 
			Category:"Makanan", 
			Stock:50,
			Description:"Mie kenyal dengan bumbu pedas manis",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"Rawon", 
			Price:25000, 
			Category:"Makanan", 
			Stock:50,
			Description:"Mawurahhhh cik",
			ImageURL:"https://picsum.photos/600/400",
		},

	}
	for _, p := range products {
	config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}