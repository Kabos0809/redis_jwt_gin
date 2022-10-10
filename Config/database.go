package Config

import (
	"github.com/joho/godotenv"
	"fmt"
	"os"
)

type DbConfig struct {
	Host string
	User string
	Pass string
	DB string
	Port string
	TZ string
}

func buildDbConfig() *DbConfig {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	dbconfig := &DbConfig{
		Host: os.Getenv("HOST"),
		User: os.Getenv("POSTGRES_USER"),
		Pass: os.Getenv("POSTGRES_PASSWORD"),
		DB: os.Getenv("POSTGRES_DB"),
		Port: "5432",
		TZ: os.Getenv("TZ"),
	}
	return dbconfig
}

func DBUrl() string {
	dbConfig := buildDbConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.DB,
		dbConfig.Port,
		dbConfig.TZ,
	)
}