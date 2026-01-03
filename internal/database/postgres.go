package database

import (
	"authify/pkg/logger"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func ConnectDB() {
	var err error

	connectionString := os.Getenv("DATABASE_URL")
	conn, err = pgx.Connect(context.Background(), connectionString)
	if err != nil {
		panic(err)
	}

	logger.Green("DB Connected")
}

func DisconnectDB() {
	if conn != nil {
		conn.Close(context.Background())
		logger.Red("DB Disconnected")
		return
	}

	logger.Yellow("DB client is empty")
}
