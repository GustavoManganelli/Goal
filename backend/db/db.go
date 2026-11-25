package db

import (
	"fmt"
	"log"
	"os"

	"example.com/goal/backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{NoLowerCase: true},
	})

	if err != nil {
		log.Fatal("Falha ao conectar ao banco:", err)
	}

	fmt.Println("Conectado ao PostgreSQL!")
}

func RunMigrations() {
	err := DB.AutoMigrate(&models.Accounts{}, &models.Banks{}, &models.GroupExpenses{}, &models.GroupMembers{}, &models.Groups{}, &models.Payments{}, &models.PixKeys{}, &models.UserDebt{}, &models.Users{})
	if err != nil {
		log.Fatal("Falha ao executar as migrações:", err)
	}

	fmt.Println("Migrações executadas com sucesso!")
}