package main

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/itmrchow/microservice-gateway/delivery/handlers"
)

const (
	public_port   = "8080" //TODO: 移到config
	internal_port = "8081" //TODO: 移到config
)

func main() {
	initLog()
	initRouter()
}

// initRouter 初始化路由
func initRouter() {
	mux := handlers.RegisterPublicHandlers()
	log.Info().Msg("http internal server listen in port " + public_port)
	log.Fatal().AnErr("error", http.ListenAndServe(":"+public_port, mux))
}

// initLog 初始化log
func initLog() {
	now := time.Now()
	zerolog.TimestampFunc = now.UTC
	// TODO: 初始化時間
	// TODO: 輸出位置
	// TODO: Global log level
}
