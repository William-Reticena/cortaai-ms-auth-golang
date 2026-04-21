package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("DATABASE_URL não configurado")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco: %v", err)
	}

	DB = db
	log.Println("✅ Conectado ao banco de dados com sucesso!")

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
