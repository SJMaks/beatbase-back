package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjmaks/beatbase-back/internal/core/config"
	"github.com/sjmaks/beatbase-back/internal/core/database"
	"github.com/sjmaks/beatbase-back/internal/core/logger"
)

func main() {
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

	engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Server healthy and running")
	})

	if err := engine.Run("127.0.0.1:8080"); err != nil {
		logger.Error(fmt.Sprintf("failed to run server: %v", err))
		log.Fatalf("failed to run server: %v", err)
	}
}
