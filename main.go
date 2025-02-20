package main

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/itmrchow/microservice-gateway/delivery/handlers"
)

func main() {
	initConfig()
	initLog()
	initRouter()
}

// initConfig 初始化config
func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("config init error")
	}
}

// initLog 初始化log
func initLog() {
	now := time.Now()
	zerolog.TimestampFunc = now.UTC
	// TODO: 初始化時間
	// TODO: 輸出位置
	// TODO: Global log level
}

// initRouter 初始化路由
func initRouter() {
	var (
		// publicPort   = viper.GetString("http_public_port")
		internalPort = viper.GetString("http_internal_port")
	)

	mux := handlers.RegisterPublicHandlers()
	log.Info().Msg("http internal server listen in port " + internalPort)
	log.Fatal().AnErr("error", http.ListenAndServe(":"+internalPort, mux))
}
