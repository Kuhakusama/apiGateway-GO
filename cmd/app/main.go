package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // Import the godotenv package
	"github.com/kuhakusama/apiGateway-GO/internal/repository"
	"github.com/kuhakusama/apiGateway-GO/internal/service"
	"github.com/kuhakusama/apiGateway-GO/internal/web/server"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != ""{
		return value
	}
	return defaultValue
}

//toda aplicação precisa de um entry point
func main(){
	if err := godotenv.Load(); err != nil { //carrega as variaveis de ambiente do arquivo .env
		log.Fatal("Error loading .env file")
	}
	
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db) //cria o repositorio de contas
	accountService := service.NewAccountService(accountRepository) //cria o servico de contas

	port := getEnv("HTTP_PORT", "8080")
	srv := server.NewServer(accountService, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}