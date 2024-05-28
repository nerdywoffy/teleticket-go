//go:generate go run github.com/steebchen/prisma-client-go db push --schema ./../../db/schema/schema.prisma

package main

import (
	"os"

	"github.com/joho/godotenv"
	db "github.com/nerdywoffy/teleticket-go/db"
	log "github.com/nerdywoffy/teleticket-go/internal/logger"
)

func main() {
	// Only load .env if file exists
	if stat, err := os.Stat(".env"); err == nil && !stat.IsDir() {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	logger := log.NewLogger()
	client := db.NewClient(db.WithDatasourceURL(os.Getenv("DB_DATASOURCE_URL")))
	if err := client.Prisma.Connect(); err != nil {
		logger.Panic(err)
	}
}
