package main

import (
	"log"

	"github.com/RIOKOWI/tugas_week_5_1123150059/config"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment variable sistem")
	}

	config.InitFirebase()
	
	config.InitDatabase()
}
