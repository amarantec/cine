package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/amarantec/cine/internal/database"
	"gitlab.com/amarantec/cine/internal/handlers"
)

func main() {
	loadEnv()
	ctx := context.Background()

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("DB_PORT")
	serverPort := os.Getenv("SERVER_PORT")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || serverPort == "" {
		log.Fatal("one or more environment variables are not set")
	}

	connectionString := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		dbHost, dbPort, dbUser, dbPassword, dbName)

	Conn, err := database.OpenConnection(ctx, connectionString)
	if err != nil {
		panic(err)
	}

	defer Conn.Close()

	mux := handlers.SetRoutes(Conn)

	port := fmt.Sprintf(":%s", serverPort)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	fmt.Printf("Server listen on %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())

}

func loadEnv() {
	err := godotenv.Load("../config/.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
