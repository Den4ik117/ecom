package main

import (
	"database/sql"
	"github.com/Den4ik117/ecom/cmd/api"
	"github.com/Den4ik117/ecom/config"
	"github.com/Den4ik117/ecom/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	database, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	initStorage(database)

	server := api.NewApiServer(":8080", database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %s", err)
	}

	log.Println("Connected to database")
}
