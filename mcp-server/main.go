package main

import (
	"context"
	"log"
	"mcp-server/src/databases"
	"mcp-server/src/register"

	"github.com/joho/godotenv"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "to-do-list", Version: "v1.0.0"}, nil)

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}

	db, err := databases.ConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}

	register.RegisterTodoTools(server, db)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
