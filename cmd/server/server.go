package main

import (
	"backendServer/internal/configs"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// load config
	cfg := configs.LoadConfig()

	// init logger
	initLogger(cfg)

	// set gin mode
	gin.SetMode(cfg.GinMode)

	router := gin.Default()

	// Определяем GET-эндпоинт
	router.GET("/ping", func(c *gin.Context) {
		// Отвечаем JSON'ом
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 6. Сервер с graceful shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// 7. Ожидание сигнала
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exited properly")
}

func initLogger(cfg *configs.AppConfig) {
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal("Cannot create logs directory:", err)
	}
	// Настройка ротации логов
	logger := &lumberjack.Logger{
		Filename:   "logs/gin.log",    // path to file
		MaxSize:    cfg.LogMaxSize,    // MB before rotation
		MaxBackups: cfg.LogMaxBackups, // old files count
		MaxAge:     cfg.LogMaxAge,     // days to store
		Compress:   cfg.LogCompress,   // compress old files
	}

	multiWriter := io.MultiWriter(logger, os.Stdout)
	gin.DefaultWriter = multiWriter
	// write logs to file and console
	log.SetOutput(multiWriter)
}
