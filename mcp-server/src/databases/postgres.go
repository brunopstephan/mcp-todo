package databases

import (
	"fmt"
	"log"
	"mcp-server/src/schemas"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {
	log.Println(os.Getenv("POSTGRES_DSN"))
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_DSN")), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting in postgres: %v", err)
	}

	err = db.AutoMigrate(&schemas.Todo{})
	if err != nil {
		return nil, fmt.Errorf("error migrating tables: %v", err)
	}

	return db, nil
}
