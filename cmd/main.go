package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nawafilhusnul/music-catalog/internal/configs"
	membershipshandler "github.com/nawafilhusnul/music-catalog/internal/handlers/memberships"
	trackshandler "github.com/nawafilhusnul/music-catalog/internal/handlers/tracks"
	"github.com/nawafilhusnul/music-catalog/internal/models/memberships"
	membershipsrepository "github.com/nawafilhusnul/music-catalog/internal/repository/memberships"
	spotifyoutbound "github.com/nawafilhusnul/music-catalog/internal/repository/spotify"
	membershipservice "github.com/nawafilhusnul/music-catalog/internal/service/memberships"
	tracksservice "github.com/nawafilhusnul/music-catalog/internal/service/tracks"
	"github.com/nawafilhusnul/music-catalog/pkg/httpclient"
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %+v\n", err)
	}

	db.AutoMigrate(&memberships.User{})

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipsRepo := membershipsrepository.NewRepository(db)
	membershipsSvc := membershipservice.NewService(cfg, membershipsRepo)
	membershipsHandler := membershipshandler.NewHandler(membershipsSvc, r)
	membershipsHandler.RegisterRoutes()

	httpClient := httpclient.NewClient(&http.Client{})
	spotifyOutbound := spotifyoutbound.NewSpotifyOutbound(cfg, httpClient)
	tracksSvc := tracksservice.NewService(spotifyOutbound)
	tracksHandler := trackshandler.NewHandler(tracksSvc, r)
	tracksHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
}
