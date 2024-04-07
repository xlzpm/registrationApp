package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xlzpm/internal/config"
	"github.com/xlzpm/internal/users"
	"github.com/xlzpm/internal/users/db/pg"
	"github.com/xlzpm/pkg/logger/initlog"
	"github.com/xlzpm/pkg/pgclient"
)

func main() {
	cfg := config.MustConfig()

	log := initlog.InitLogger()

	pgClient, err := pgclient.NewPostgresDB(context.TODO(), cfg.Storage)
	if err != nil {
		log.Error(err.Error())
		pgClient.Close()
	}

	repo := pg.NewRepository(pgClient)
	log.Info("register users handler and users route")
	userHandler := users.NewHandler(repo)
	userroute := userHandler.InitRoutes()

	start(userroute, cfg)
}

func start(router *gin.Engine, cfg *config.Config) {
	log := initlog.InitLogger()

	log.Info("listen tcp")

	var listener net.Listener
	var listenErr error

	listener, listenErr = net.Listen("tcp",
		fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
	log.Info("server is listening port %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)

	if listenErr != nil {
		log.Error(listenErr.Error())
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Error(server.Serve(listener).Error())
}
