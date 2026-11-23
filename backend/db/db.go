package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	// Access environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Falha ao conectar ao banco:", err)
	}

	fmt.Println("Conectado ao PostgreSQL!")
}

func RunMigrations() {
	err := DB.AutoMigrate()
	if err != nil {
		log.Fatal("Falha ao executar as migrações:", err)
	}

	fmt.Println("Migrações executadas com sucesso!")
}