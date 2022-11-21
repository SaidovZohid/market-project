package main

import (
	"fmt"
	"log"

	"github.com/SaidovZohid/market-project/api"
	"github.com/SaidovZohid/market-project/config"
	"github.com/SaidovZohid/market-project/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load(".")

	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psql)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	strg := storage.NewStoragePg(psqlConn)

	api := api.New(&api.RouteOptions{
		Cfg: &cfg,
		Storage: strg,
	})

	api.Run(cfg.HttpPort)
}
