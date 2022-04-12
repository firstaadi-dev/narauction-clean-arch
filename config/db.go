package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	godotenv.Load()
	DSN := os.Getenv("DSN")
	db, _ := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	//db.AutoMigrate(&domain.Barang{})
	//db.AutoMigrate(&domain.Event{})
	return db
}
