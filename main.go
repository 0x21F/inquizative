package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/0x21F/inquizative/routes"
	"github.com/0x21F/inquizative/types"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("bruh")
	}

	url := os.Getenv("DB_URL") + "?authToken=" + os.Getenv("DB_TOKEN")
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal(err)
	}

	err = types.UpTables(db)
	if err != nil {
		log.Fatal(err)
	}

	controller := routes.RouteController{}
	e := controller.ToServer(db)

	e.Start(":8080")
}
