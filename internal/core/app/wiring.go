package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjmaks/beatbase-back/internal/core/config"
	"github.com/sjmaks/beatbase-back/internal/core/database"
	"github.com/sjmaks/beatbase-back/internal/core/logger"
	httpserver "github.com/sjmaks/beatbase-back/internal/core/transport/http"
)

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.MustLoad()

	logger, closeLog := logger.MustNewZap()
	defer func() {
		if err := closeLog(); err != nil {
			log.Fatalf("failed to close log file: %v", err)
		}
	}()

	pool := database.MustNewPool(ctx, logger, cfg)
	defer pool.Close()

	engine := gin.New()

	engine.GET("/healthz", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	srv := httpserver.NewHTTPServer(":8080", engine)

	err := srv.ListenAndServe()
	if err != nil {
		logger.Error(fmt.Sprintf("failed to run server: %v", err))
		log.Fatalf("failed to run server: %v", err)
	}
	return err
}
