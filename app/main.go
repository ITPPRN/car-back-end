package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"testBackend/configs"
	"testBackend/modules/servers"
	databases "testBackend/pkg/postgres"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := new(configs.Config)

	//App env
	cfg.App.Port = os.Getenv("APP_PORT")
	cfg.App.AllowOrigins = os.Getenv("CORS_ALLOW_ORIGIN")

	//postgres env
	cfg.Postgres.Host = os.Getenv("DB_HOST")
	cfg.Postgres.Port = os.Getenv("DB_PORT")
	cfg.Postgres.Username = os.Getenv("DB_USER")
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")
	cfg.Postgres.DatabaseName = os.Getenv("DB_NAME")
	cfg.Postgres.SslMode = os.Getenv("DB_SSLMODE")
	cfg.Postgres.Schema = os.Getenv("DB_SCHEMA")

	db, err := databases.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	server := servers.NewServer(cfg, db)
	server.Start()
}
