package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var DB *gorm.DB

var AppConfig *Config

type Config struct {
	AppPort string
	
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret string
	JWTExpirationMinutes string
	JWTRefreshToken string
	JWTExpiration string
}

func LoadEnv(){
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using default values")
	}

	AppConfig = &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBName: getEnv("DB_NAME", "phonebook_db"),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
		JWTExpirationMinutes: getEnv("JWT_EXPIRATION_MINUTES", "15"),
		JWTRefreshToken: getEnv("REFRESH_TOKEN_EXPIRATION", "refresh_token"),
		JWTExpiration: getEnv("JWT_EXPIRATION", "1h"),
	}
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}
	return value
}

func ConnectDB(){
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		AppConfig.DBHost,
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBName,
		AppConfig.AppPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic("failed to connect database")
	}

	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
}