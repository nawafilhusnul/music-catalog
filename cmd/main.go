package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nawafilhusnul/music-catalog/internal/configs"
	"github.com/nawafilhusnul/music-catalog/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolders([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("failed to initialize config: %v\n", err)
	}

	cfg = configs.Get()

	_, err = internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %+v\n", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run(cfg.Service.Port)
}
