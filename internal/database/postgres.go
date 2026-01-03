package database

import (
	"authify/pkg/logger"
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	connectionString := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		panic(err)
	}

	logger.Green("DB Connected")

	defer func() {
		conn.Close(context.Background())
		logger.Red("DB Disconnected")
	}()
}
