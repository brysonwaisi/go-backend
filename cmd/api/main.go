package main

import (
	"log"

	"github.com/brysonwaisi/go-backend/internal/env"
	"github.com/brysonwaisi/go-backend/internal/store"

)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db dbConfig {
			addr: env.GetString("DB_ADDR", "postgres://user:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime: env.GetInt("DB_MAX_IDLE_TIME", "15min")
		},
	}

	store := store.NewStorage(nil)

	app := &application {
		config: cfg,
		store: store,
	}


	mux :=app.mount() 

	log.Fatal(app.run(mux))
}