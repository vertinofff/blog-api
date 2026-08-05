package main

import (
	"github.com/vertinofff/blog-api/api"
	"github.com/vertinofff/blog-api/config"
	"github.com/vertinofff/blog-api/data/db"
	"github.com/vertinofff/blog-api/data/migration"
	"github.com/vertinofff/blog-api/pkg/logging"
)

func main(){
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := db.InitDb(cfg)
	defer db.CloseDb()

	if err != nil {
		logger.Fatal(logging.Postgres , logging.Startup , err.Error(),nil)
	}

	migration.Up()
	api.InitServer()
}