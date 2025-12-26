package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pauloappbr/go-vue-crm/backend/internal/config"
	"github.com/pauloappbr/go-vue-crm/backend/internal/logger"
)

func main() {
	logger.InitLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Falha ao carregar configurações", "error", err)
		os.Exit(1)
	}

	slog.Info("Configurações carregadas com sucesso", "port", cfg.Port)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		slog.Info("Request recebido", "path", "/ping", "method", "GET")
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	slog.Info("Servidor iniciando...", "port", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		slog.Error("Erro ao rodar servidor", "error", err)
	}
}
