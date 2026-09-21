package main

import (
	"context"
	"github.com/vertinofff/blog-api/api"
	"github.com/vertinofff/blog-api/config"
	"github.com/vertinofff/blog-api/data/db"
	"github.com/vertinofff/blog-api/data/migration"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, e := config.Load()
	if e != nil {
		log.Fatal(e)
	}
	database, e := db.Open(cfg)
	if e != nil {
		log.Fatal(e)
	}
	defer func() {
		if e := db.Close(database); e != nil {
			log.Printf("close database: %v", e)
		}
	}()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if os.Getenv("RUN_MIGRATIONS") == "true" {
		if e := migration.Up(ctx, database); e != nil {
			log.Fatal(e)
		}
	}
	if e := api.Run(ctx, api.NewServer(cfg, database)); e != nil {
		log.Fatal(e)
	}
}
