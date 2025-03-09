package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jamalfox85/loyalty-rewards-api/data"
	"github.com/joho/godotenv" // Import the godotenv package
)

func NewApplication() *Application {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create DB Instance and table repositories
	db := newDB()
	// defer db.Close(context.Background())

	plans := data.NewPlanRepository(db)


	return &Application{
		Plans: plans,
	}
}

func newDB() *pgx.Conn {
	connString := os.Getenv("DB_CONNECTION_STRING")

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err);
		// log.Fatal("Error initializing database");
	}

	fmt.Println("Connected to DBS!")
	return db;
}